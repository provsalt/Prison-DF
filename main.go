package main

import (
	"Prison/economy"
	"Prison/prisons/commands"
	"Prison/prisons/console"
	"Prison/prisons/handlers"
	"Prison/prisons/handlers/worlds"
	"Prison/prisons/tasks/broadcast"
	"Prison/prisons/tasks/minereset"
	"Prison/prisons/tasks/restart"
	"Prison/prisons/utils"
	"Prison/ranks"
	"fmt"
	"github.com/bradhe/stopwatch"
	_ "github.com/davecgh/go-spew/spew"
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/item"
	"github.com/df-mc/dragonfly/dragonfly/item/tool"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/player/chat"
	"github.com/df-mc/dragonfly/dragonfly/player/scoreboard"
	"github.com/df-mc/dragonfly/dragonfly/player/title"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/df-mc/dragonfly/dragonfly/world/gamemode"
	worldmanager "github.com/emperials/df-worldmanager"
	"github.com/go-resty/resty/v2"
	"github.com/nakabonne/gosivy/agent"
	"github.com/pelletier/go-toml"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	watch := stopwatch.Start()
	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "15:04:05",
			LogFormat:       "[%time%] [Server thread/%lvl%]: %msg% \n",
		},
	}
	log.Level = logrus.DebugLevel
	chat.Global.Subscribe(chat.StdoutSubscriber{})

	config, err := ReadConfig()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}
	Server := dragonfly.New(&config, log)

	err = agent.Listen(agent.Options{
		Addr: ":25565",
	})
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Println(http.ListenAndServe("localhost:5000", nil))
	}()

	if err = Server.Start(); err != nil {
		log.Fatalln(err)
	}
	log.Infof(text.ANSI(text.Colourf("<green>Starting world</green>")))
	manager, err := startWorld(Server, log)
	if err != nil {
		log.Panic(err)
	}

	log.Infof(text.ANSI(text.Colourf("<green>Registering commands</green?")))
	register := commands.Register()
	if register {
		log.Info(text.ANSI(text.Colourf("<green>Successfully registered commands</green>")))
	}

	log.Infof(text.ANSI(text.Colourf("<green>Registering databases</green?")))
	e, r := startDB(log)
	log.Infof(text.ANSI(text.Colourf("<green>Registered databases</green?")))

	utils.Server = Server
	utils.Logger = log
	utils.Worldmanager = manager
	utils.Economy = &e
	utils.Ranks = &r

	log.Infof(text.ANSI(text.Colourf("<green>Registering tasks</green>")))
	go broadcast.New()
	go func() {
		for range time.Tick(time.Minute) {
			restart.New()
		}
	}()

	minereset.NewResetAll()
	log.Infof(text.ANSI(text.Colourf("<green>Registered tasks</green>")))

	log.Infof("If you find this project useful, please consider donating to support development: " + text.ANSI(text.Colourf("<aqua>https://www.patreon.com/sandertv</aqua>")))
	watch.Stop()
	log.Infof("Done loading server in %dms", watch.Milliseconds())
	console.StartConsole()

	stop := make(chan os.Signal, 2)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-stop
		func() {
			log.Infof("Shutting down any other existing processes")
			_ = manager.Close()
			if err = Server.Close(); err != nil {
				log.Errorf("error shutting down server: %v", err)
			}
			utils.Economy.Close()
			agent.Close()
		}()
	}()
	for {
		p, err := Server.Accept()
		if err != nil {
			break
		}
		onJoin(p)
	}
	log.Infof("Shutting down any other existing processes")
	utils.Economy.Close()
	_ = manager.Close()
	agent.Close()
}

// ReadConfig reads the configuration from the config.toml file, or creates the file if it does not yet exist.
func ReadConfig() (dragonfly.Config, error) {
	c := dragonfly.DefaultConfig()
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return c, fmt.Errorf("failed encoding default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return c, fmt.Errorf("failed creating config: %v", err)
		}
		return c, nil
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return c, fmt.Errorf("error reading config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return c, fmt.Errorf("error decoding config: %v", err)
	}
	return c, nil
}
func onJoin(p *player.Player) {
	go utils.Economy.InitPlayer(p, 2000)
	p.SetGameMode(gamemode.Survival{})
	rest := resty.New()
	type json struct {
		Username string `json:"username"`
		Content  string `json:"content"`
	}
	_, err := rest.R().SetBody(json{"Joins", "[+] " + p.Name()}).Post(utils.WebhookURL)

	if err != nil {
		utils.Logger.Errorln(err)
	}
	p.Handle(handlers.NewSpawmHandler(p))
	p.ShowCoordinates()
	p.Inventory().AddItem(item.NewStack(item.Pickaxe{Tier: tool.TierIron}, 1))
	t := title.New(utils.GetPrefix())
	t = t.WithSubtitle(text.Colourf("<aqua>Season 1</aqua>"))
	time.AfterFunc(time.Second*2, func() {
		utils.Session_writePacket(utils.Player_session(p), &packet.ActorEvent{
			EventType:       packet.ActorEventElderGuardianCurse,
			EntityRuntimeID: 1,
		})
		p.SendTitle(t.WithFadeOutDuration(time.Second * 7))
	})
	time.AfterFunc(time.Second, func() {
		SendScoreBoard(p)
	})
}

func startWorld(server *dragonfly.Server, logger *logrus.Logger) (*worldmanager.WorldManager, error) {
	w := server.World()
	w.SetDefaultGameMode(gamemode.Survival{})
	w.SetSpawn(world.BlockPos{173, 98, 131})
	w.Handle(worlds.NewSpawnWorldHandler(w))

	dir, _ := os.Getwd()
	manager := worldmanager.New(server, dir, logger)

	err := manager.LoadWorld("worlds/mine_a", "mine_a", 4)

	if err != nil {
		return nil, err
	}
	return manager, nil
}

func startDB(log *logrus.Logger) (economy.Economy, ranks.RankApi) {
	e := economy.New(economy.Connection{
		Username: "u3426_jj0S8B64nZ",
		Password: "ecW=3pMEkIZkHzrPm5@zA4cU",
		IP:       "140.82.11.202:3306",
		Schema:   "s3426_prisons",
	}, 2, 10)
	r := ranks.New(ranks.Connection{
		IP:       "140.82.11.202:3306",
		Username: "u3426_jj0S8B64nZ",
		Password: "ecW=3pMEkIZkHzrPm5@zA4cU",
		Schema:   "s3426_prisons",
	}, 2, 10, log)
	return e, r
}

func SendScoreBoard(player *player.Player) {
	err, bal := utils.Economy.Balance(player)
	if err != nil {
		player.Disconnect(text.Colourf(utils.GetPrefix() + "An error occured. Please contact the staff team."))
		utils.GetLogger().Errorf("This error is caused by sebding a scoreboard: \n %w", err)
	}
	s := scoreboard.New(text.Colourf(utils.GetPrefix() + "<aqua><b>Prisons</b></aqua>")).Addf(text.Colourf("<b><dark-grey>*</dark-grey><gold>%s</gold><red>%v</red></b>", "Your balance: ", bal))
	player.SendScoreboard(s)
}

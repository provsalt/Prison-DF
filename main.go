package main

import (
	"Prison/economy"
	"Prison/prisons/commands"
	"Prison/prisons/console"
	"Prison/prisons/handlers"
	"Prison/prisons/tasks/broadcast"
	"Prison/prisons/tasks/restart"
	"Prison/prisons/utils"
	"fmt"
	"github.com/bradhe/stopwatch"
	_ "github.com/davecgh/go-spew/spew"
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/player/chat"
	"github.com/df-mc/dragonfly/dragonfly/player/title"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/df-mc/dragonfly/dragonfly/world/gamemode"
	worldmanager "github.com/emperials/df-worldmanager"
	"github.com/nakabonne/gosivy/agent"
	"github.com/pelletier/go-toml"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {
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
	watch := stopwatch.Start()
	Server := dragonfly.New(&config, log)

	defer agent.Close()
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
		return
	}
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	Server.CloseOnProgramEnd()
	if err := Server.Start(); err != nil {
		log.Fatalln(err)
	}
	log.Infof(text.ANSI(text.Colourf("<green>Starting world</green>")))
	w := Server.World()
	w.SetDefaultGameMode(gamemode.Survival{})
	w.SetSpawn(world.BlockPos{0, 4, 0})
	w.SetSpawn([3]int{173, 98, 131})

	dir, _ := os.Getwd()
	manager := worldmanager.New(Server, dir, log)

	err = manager.LoadWorld("worlds/mine_a", "mine_a", 4)

	if err != nil {
		panic(err)
	}

	console.StartConsole()

	log.Infof(text.ANSI(text.Colourf("<green>Registering commands</green?")))
	register := commands.Register()
	if register {
		log.Info(text.ANSI(text.Colourf("<green>Successfully registered commands</green>")))
	}

	e := economy.New(economy.Connection{
		Username: "u1990_9jqSt4O0ET",
		Password: "2z^lICvFF86g^sW5Lcp=tc6E",
		IP:       "140.82.11.202:3306",
		Schema:   "s1990_economy",
	}, 3, 10)

	utils.Server = Server
	utils.Logger = log
	utils.Worldmanager = manager
	utils.Economy = &e

	log.Infof(text.ANSI(text.Colourf("<green>Registering tasks</green>")))
	go broadcast.New()
	go func() {
		for range time.Tick(time.Minute) {
			restart.New()
		}
	}()
	log.Infof(text.ANSI(text.Colourf("<green>Registered tasks</green>")))

	log.Infof("If you find this project useful, please consider donating to support development: " + text.ANSI(text.Colourf("<aqua>https://www.patreon.com/sandertv</aqua>")))
	watch.Stop()
	log.Infof("Done loading server in %dms", watch.Milliseconds())
	for {
		p, err := Server.Accept()
		if err != nil {
			break
		}
		onJoin(p)
	}
	err = manager.Close()
	if err != nil {
		panic(err)
	}
	utils.Economy.Close()
}

// ReadConfig reads the configuration from the config.toml file, or creates the file if it does not yet exist.
func ReadConfig() (dragonfly.Config, error) {
	c := dragonfly.DefaultConfig()
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return c, fmt.Errorf("failed encoding default config: %v", err)
		}
		if err := ioutil.WriteFile("config.toml", data, 0644); err != nil {
			return c, fmt.Errorf("failed creating config: %v", err)
		}
		return c, nil
	}
	data, err := ioutil.ReadFile("config.toml")
	if err != nil {
		return c, fmt.Errorf("error reading config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return c, fmt.Errorf("error decoding config: %v", err)
	}
	return c, nil
}
func onJoin(p *player.Player) {
	p.SetGameMode(gamemode.Survival{})
	p.Handle(handlers.NewSpawmHandler(p))
	p.ShowCoordinates()
	t := title.New(utils.GetPrefix())
	t = t.WithSubtitle(text.Colourf("<aqua>Season 1</aqua>"))
	time.AfterFunc(time.Second*3, func() {
		utils.Session_writePacket(utils.Player_session(p), &packet.ActorEvent{
			EventType:       packet.ActorEventElderGuardianCurse,
			EntityRuntimeID: 1,
		})
		p.SendTitle(t.WithFadeOutDuration(time.Second * 7))
	})
	utils.Economy.InitPlayer(p, 2000)
}

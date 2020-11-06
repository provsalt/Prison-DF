package main

import (
	"Prison/prisons/commands"
	"Prison/prisons/console"
	"Prison/prisons/events"
	"Prison/prisons/tasks/broadcast"
	"Prison/prisons/utils"
	"fmt"
	"github.com/bradhe/stopwatch"
	_ "github.com/davecgh/go-spew/spew"
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/player/chat"
	"github.com/df-mc/dragonfly/dragonfly/player/title"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/df-mc/dragonfly/dragonfly/world/gamemode"
	worldmanager "github.com/emperials/df-worldmanager"
	"github.com/pelletier/go-toml"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io/ioutil"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3)
	fmt.Println(runtime.GOMAXPROCS(0))
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
	Server.CloseOnProgramEnd()
	if err := Server.Start(); err != nil {
		log.Fatalln(err)
	}
	log.Infof(text.ANSI(text.Colourf("<green>Starting world</green>")))
	w := Server.World()
	w.SetDefaultGameMode(gamemode.Survival{})
	w.SetSpawn(world.BlockPos{0, 4, 0})
	w.SetTime(5000)
	w.StopTime()

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

	utils.Server = Server
	utils.Logger = log
	utils.Worldmanager = manager

	log.Infof(text.ANSI(text.Colourf("<green>Registering tasks</green>")))
	go broadcast.New()
	log.Infof(text.ANSI(text.Colourf("<green>Registered tasks</green>")))

	watch.Stop()
	log.Infof("Done loading server in %dms", watch.Milliseconds())

	for {
		p, err := Server.Accept()
		if err != nil {
			break
		}
		p.Handle(events.NewPlayerQuitHandler(p))
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
		p.SetGameMode(gamemode.Creative{})
		ws, _ := manager.World("mine_a")
		ws.AddEntity(p)
		p.Teleport(ws.Spawn().Vec3Middle())
	}
	err = manager.Close()

	if err != nil {
		panic(err)
	}
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

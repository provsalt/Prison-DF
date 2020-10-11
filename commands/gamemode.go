package commands

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/world/gamemode"
	"strings"
)

type Gamemode struct {
	Gamemode string
}

func (g Gamemode) Run(source cmd.Source, output *cmd.Output) {
	if _, ok := source.(*player.Player); !ok {
		output.Printf("You must run this command as a player")
	}

	p := source.(*player.Player)
	mode := strings.ToLower(g.Gamemode)
	switch mode {
	case "creative", "c":
		p.SetGameMode(gamemode.Creative{})
		output.Printf("Set your own gamemode to creative")
	case "survival", "s":
		p.SetGameMode(gamemode.Survival{})
		output.Printf("Set your own gamemode to survival")
	case "adventure", "a":
		p.SetGameMode(gamemode.Adventure{})
		output.Printf("Set your own gamemode to adventure")
	case "spectator":
		p.SetGameMode(gamemode.Spectator{})
		output.Printf("Set your own gamemode to spectator")
	default:
		output.Printf("Unknown game mode")
	}
}

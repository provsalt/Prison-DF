package gamemode

import (
	"Prison/prisons/database/ranks"
	"Prison/prisons/utils"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/world/gamemode"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"strings"
)

type Gamemode struct {
	Gamemode string
}

func (g Gamemode) Run(source cmd.Source, output *cmd.Output) {
	if _, ok := source.(*player.Player); !ok {
		output.Printf(text.ANSI(text.Colourf("<red>You must run this command as a player</red>")))
		return
	}

	p := source.(*player.Player)

	staffRank := utils.RanksDB.GetPermissionLevel(p).StaffRanks
	if staffRank < ranks.Manager {
		output.Printf(text.Colourf("Haha, nice try. You aren't suppose to run this."))
	}

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

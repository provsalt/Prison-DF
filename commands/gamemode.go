package commands

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/world/gamemode"
)

type Gamemode struct{}

func (g Gamemode) Run(source cmd.Source, output *cmd.Output) {
	p := source.(*player.Player)

	if p.Name() == "provsalt" {
		p.SetGameMode(gamemode.Creative{})
	}
}

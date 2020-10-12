package stop

import (
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
)

var Server *dragonfly.Server

type Stop struct{}

func (s Stop) Run(source cmd.Source, output *cmd.Output) {
	_, ok := source.(*player.Player)

	if !ok {
		err := Server.Close()
		if err != nil {
			panic(err)
		}
	}
}

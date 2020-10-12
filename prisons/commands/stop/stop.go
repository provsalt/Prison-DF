package stop

import (
	"Prison/prisons/utils"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
)

type Stop struct{}

func (s Stop) Run(source cmd.Source, output *cmd.Output) {
	_, ok := source.(*player.Player)

	if !ok {
		err := utils.GetServer().Close()
		if err != nil {
			panic(err)
		}
	}
}

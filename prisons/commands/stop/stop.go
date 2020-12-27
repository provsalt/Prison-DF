package stop

import (
	"Prison/prisons/utils"
	ranks2 "Prison/ranks"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
)

type Stop struct{}

func (s Stop) Run(source cmd.Source, output *cmd.Output) {
	p, ok := source.(*player.Player)

	if ok {
		ranks := utils.Ranks.GetPermissionLevel(p)
		if ranks.StaffRanks != ranks2.Owner {
			return
		}
		ok = false
	}

	if !ok {
		err := utils.GetServer().Close()
		if err != nil {
			panic(err)
		}
	}
}

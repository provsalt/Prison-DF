package stop

import (
	ranks2 "Prison/prisons/database/ranks"
	"Prison/prisons/utils"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
)

type Stop struct{}

func (s Stop) Run(source cmd.Source, output *cmd.Output) {
	p, ok := source.(*player.Player)

	if ok {
		ranks := utils.RanksDB.GetPermissionLevel(p)
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

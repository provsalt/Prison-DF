package minereset

import (
	"Prison/prisons/tasks"
	"Prison/prisons/utils"
	"github.com/df-mc/dragonfly/dragonfly/block"
	"github.com/df-mc/dragonfly/dragonfly/player/chat"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"time"
)

func NewResetAll() {
	tasks.RepeatingTask(func() {
		uptime := utils.GetServer().Uptime()
		switch {
		case uptime > time.Minute*15, uptime > time.Minute*30, uptime > time.Minute*45, uptime > time.Minute*59:
			chat.Global.Printf(text.Colourf(utils.Broadcastprefix + "<red>Reseting spawn mine\n</red>"))
			reset := MineReset{Mine: struct {
				MineName  string
				Dimension [3]int
				Blocks    map[world.Block]interface{}
			}{MineName: "spawn", Dimension: [3]int{55, 37, 128}, Blocks: map[world.Block]interface{}{
				block.Stone{}:     [2]int{0, 79},
				block.CoalOre{}:   [2]int{80, 95},
				block.CoalBlock{}: [2]int{96, 100},
			}}}
			utils.GetServer().World().BuildStructure(world.BlockPosFromVec3(mgl64.Vec3{145, 57, 218}), reset)

		}
	}, time.Minute*15)
}

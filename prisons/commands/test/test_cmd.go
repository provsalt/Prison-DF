package test

import (
	"Prison/prisons/console"
	"Prison/prisons/tasks/minereset"
	"Prison/prisons/utils"
	"github.com/bradhe/stopwatch"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type TestCmd struct{}

func (t TestCmd) Run(source cmd.Source, output *cmd.Output) {
	if _, ok := source.(*console.Console); !ok {
		return
	}
	s := stopwatch.Start()
	reset := minereset.MineReset{Mine: struct {
		MineName  string
		Dimension [3]int
		Blocks    map[world.Block][2]int
	}{MineName: "spawn", Dimension: [3]int{55, 37, 128}, Blocks: map[world.Block][2]int{
		block.Stone{}:     {0, 60},
		block.CoalOre{}:   {60, 90},
		block.CoalBlock{}: {90, 100},
	}}}
	utils.GetServer().World().BuildStructure(cube.PosFromVec3(mgl64.Vec3{145, 57, 218}), reset)
	s.Stop()
	utils.GetLogger().Infof("Done reseting in %dms", s.Milliseconds()) // Done reseting in 5278ms
}

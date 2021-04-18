package worlds

import (
	"github.com/df-mc/dragonfly/dragonfly/block/cube"
	"github.com/df-mc/dragonfly/dragonfly/event"
	"github.com/df-mc/dragonfly/dragonfly/world"
)

type SpawnWorldHandler struct {
	w *world.World
	world.NopHandler
}

func NewSpawnWorldHandler(world *world.World) *SpawnWorldHandler {
	return &SpawnWorldHandler{w: world}
}

func (s SpawnWorldHandler) HandleLiquidHarden(event *event.Context, _ cube.Pos, _ world.Block, _ world.Block, _ world.Block) {
	if s.w.Name() == "spawn" {
		event.Cancel()
	}
}
func (s SpawnWorldHandler) HandleLiquidFlow(event *event.Context, _ cube.Pos, _ cube.Pos, _ world.Block, _ world.Block) {
	if s.w.Name() == "spawn" {
		event.Cancel()
	}
}

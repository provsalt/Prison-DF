package worlds

import "github.com/df-mc/dragonfly/dragonfly/world"

type SpawnHandler struct {
	world.NopHandler
	w *world.World
}

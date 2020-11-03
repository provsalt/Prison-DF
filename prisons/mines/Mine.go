package mines

import (
	"github.com/df-mc/dragonfly/dragonfly/world"
)

type Mine struct {
	MineName string

	// Blocks followed by the block. Must total to 100
	Blocks map[world.Block]interface{}
}

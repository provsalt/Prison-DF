package mines

import (
	"github.com/df-mc/dragonfly/server/world"
)

type Mine struct {
	MineName string

	Dimension [3]int
	// Blocks followed by the block. Must total to 100
	Blocks map[world.Block][2]int
}

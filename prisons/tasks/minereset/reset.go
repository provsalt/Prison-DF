package minereset

import (
	"Prison/prisons/mines"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"math/rand"
	"time"
)

type MineReset struct {
	Mine mines.Mine
}

func (m MineReset) Dimensions() [3]int {
	return m.Mine.Dimension
}

func (m MineReset) At(x, y, z int, blockAt func(x int, y int, z int) world.Block) world.Block {
	for block, i := range m.Mine.Blocks {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		chance := r.Intn(101)
		for k, j := range i {
			if chance >= k && chance <= j {
				return block
			}
		}
	}
	return nil
}

func (m MineReset) AdditionalLiquidAt(x, y, z int) world.Liquid {
	return nil
}

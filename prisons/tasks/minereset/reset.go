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
	return [3]int{64, 64, 64}
}

func (m MineReset) At(x, y, z int, blockAt func(x int, y int, z int) world.Block) world.Block {
	for block, i := range m.Mine.Blocks {
		chance := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
		for k, j := range i.([2]int) {
			// k refers to the lowest
			// j refers to the highest
			if chance >= k && chance <= j {
				return block
			}
			continue
		}
	}
	return nil
}

func (m MineReset) AdditionalLiquidAt(x, y, z int) world.Liquid {
	return nil
}

package minereset

import (
	"Prison/prisons/mines"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"sync"
)

type MineReset struct {
	Mine mines.Mine
}

func (m MineReset) Dimensions() [3]int {
	return m.Mine.Dimension
}

func (m MineReset) At(x, y, z int, blockAt func(x int, y int, z int) world.Block) world.Block {
	// TODO BIG TODO MAKE USE OF CHANNELS AND GOROUTINES this is way too slow
	wg := sync.WaitGroup{}
	wg.Add(1)
	// TODO I can probably use len() and just rand for this instead

	// for block, i := range m.Mine.Blocks {
	// 	chance := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	// 	for k, j := range i.([2]int) {
	// 		// k refers to the lowest
	// 		// j refers to the highest
	// 		if chance >= k && chance <= j {
	// 			return block
	// 		}
	// 		continue
	// 	}
	// }
	return nil
}

func (m MineReset) AdditionalLiquidAt(x, y, z int) world.Liquid {
	return nil
}

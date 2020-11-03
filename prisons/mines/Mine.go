package mines

import "github.com/df-mc/dragonfly/dragonfly/entity/physics"

type Mine struct {
	MineName string

	Location physics.AABB
}

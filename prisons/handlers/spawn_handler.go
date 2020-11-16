package handlers

import (
	"github.com/df-mc/dragonfly/dragonfly/entity"
	"github.com/df-mc/dragonfly/dragonfly/event"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"sync"
)

type SpawnHandler struct {
	p *player.Player
	player.NopHandler
}

var handlers sync.Map

func NewSpawmHandler(player *player.Player) *SpawnHandler {
	h := &SpawnHandler{
		p: player,
	}
	handlers.Store(player, h)
	return h

}

func (handler SpawnHandler) HandleQuit() {
	// Storage is next.
}

func (handler SpawnHandler) HandleItemDrop(event *event.Context, item *entity.Item) {
	event.Continue(func() {
		if item.World().Name() == "spawn" {
			event.Cancel()
		}
	})
}

func (handler SpawnHandler) HandleAttackEntity(event *event.Context, entity world.Entity) {
	if _, ok := entity.(*player.Player); ok {
		event.Cancel()
	}
}

package handlers

import (
	"github.com/df-mc/dragonfly/dragonfly/entity"
	"github.com/df-mc/dragonfly/dragonfly/event"
	"github.com/df-mc/dragonfly/dragonfly/player"
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

}

package events

import (
	"github.com/df-mc/dragonfly/dragonfly/player"
	"sync"
)

type PlayerQuit struct {
	p *player.Player
	player.NopHandler
}

var handlers sync.Map

func NewPlayerQuitHandler(player *player.Player) *PlayerQuit {
	h := &PlayerQuit{
		p: player,
	}
	handlers.Store(player, h)
	return h
}

func (receiver PlayerQuit) HandleQuit() {
	// Storage is next.
}

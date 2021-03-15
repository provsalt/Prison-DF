package handlers

import (
	"Prison/prisons/utils"
	"Prison/ranks"
	"strings"
	"sync"

	"github.com/df-mc/dragonfly/dragonfly/entity/physics"
	"github.com/df-mc/dragonfly/dragonfly/event"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/player/chat"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/go-resty/resty/v2"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type SpawnHandler struct {
	p *player.Player
	player.NopHandler
	ranks ranks.Ranks
}

var handlers sync.Map

func NewSpawmHandler(player *player.Player) *SpawnHandler {
	go utils.Ranks.InitPlayer(player)
	ranks2 := utils.Ranks.GetPermissionLevel(player)
	h := &SpawnHandler{
		p:     player,
		ranks: ranks2,
	}
	handlers.Store(player, h)
	return h

}

func (h SpawnHandler) HandleQuit() {
	// TODO: Storage is next.
	if utils.Development {
		return
	}
	type json struct {
		Username string `json:"username"`
		Content  string `json:"content"`
	}
	rest := resty.New()
	_, err := rest.R().SetBody(json{"Leaves", "[-] " + h.p.Name()}).Post(utils.WebhookURL)

	if err != nil {
		utils.Logger.Errorln(err)
	}
}

func (handler SpawnHandler) HandleBlockBreak(event *event.Context, pos world.BlockPos) {
	if handler.p.World().Name() == "spawn" {
		spawn := physics.NewAABB(mgl64.Vec3{145, 57, 218}, mgl64.Vec3{201, 95, 274})
		if !spawn.Vec3Within(pos.Vec3()) {
			handler.p.SendTip(text.Colourf("<red>You are not allowed to break blocks here</red>"))
			event.Cancel()
		}
	}
}

func (handler SpawnHandler) HandleBlockPlace(event *event.Context, pos world.BlockPos, block world.Block) {
	if handler.p.World().Name() == "spawn" {
		handler.p.SendTip(text.Colourf("<red>You are not allowed to place blocks here</red>"))
		event.Cancel()
	}
}

func (h SpawnHandler) HandleChat(event *event.Context, msg *string) {
	message := strings.Builder{}

	switch h.ranks.StaffRanks {
	case ranks.Owner:
		message.WriteString(text.Colourf("<b><red>[<green>OWNER</green>]</red></b> "))
	case ranks.Manager:
		message.WriteString(text.Colourf("<b><dark-yellow>[MANAGER]</dark-yellow</b> "))
	case ranks.Moderator:
		message.WriteString(text.Colourf("<dark-green>[MODERATOR]</dark-green> "))
	case ranks.Helper:
		message.WriteString(text.Colourf("<dark-blue>[HELPER]</dark-blue> "))
	}

	switch h.ranks.PaidRanks {
	case ranks.Coal:
		message.WriteString(text.Colourf("<grey>[Coal}</grey> "))
	case ranks.Gold:
		message.WriteString(text.Colourf("<gold>[Gold]</gold> "))
	case ranks.Diamond:
		message.WriteString(text.Colourf("<b><aqua>[Diamond]</aqua><b> "))
	case ranks.Emerald:
		message.WriteString(text.Colourf("<b><green>[EMERALD]</green><b> "))
	case ranks.Netherite:
		message.WriteString(text.Colourf("<b><black>[<red>NETHE</red><dark-grey>RITE</dark-grey>]</black><b> "))
	}

	for r, i := range ranks.GetAllPrisonRanks() {
		if i == h.ranks.PrisonRanks {
			message.WriteString(text.Colourf("<grey>[</grey><green>%s</green><grey>]</grey> ", r))
		}
	}
	message.WriteString(h.p.Name() + ": " + *msg)
	event.Cancel()
	chat.Global.Println(message.String())
	if utils.Development {
		return
	}

	type json struct {
		Username string `json:"username"`
		Content  string `json:"content"`
	}
	rest := resty.New()
	_, err := rest.R().SetBody(json{"Prisons Chat", h.p.Name() + ": " + *msg}).Post(utils.WebhookURL)

	if err != nil {
		utils.Logger.Errorln(err)
	}
}

func (h SpawnHandler) HandleMoneyChange(ctx event.Context, bal int) {

}

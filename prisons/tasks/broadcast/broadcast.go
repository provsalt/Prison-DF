package broadcast

import (
	"Prison/prisons/utils"

	"github.com/df-mc/dragonfly/server/player/chat"

	"time"

	"github.com/sandertv/gophertunnel/minecraft/text"
)

var Messages = [...]string{
	text.Colourf(utils.Broadcastprefix + "Join our discord today at discord.moonlightpe.com\n"),
	text.Colourf(utils.Broadcastprefix + "Remember to vote for our server at vote.moonlightpe.com for rewards\n"),
	text.Colourf(utils.Broadcastprefix + "Follow our twitter @RealMoonlightPE\n"),
	text.Colourf(utils.Broadcastprefix + "Check dragonfly out, our server software at github.com/df-mc/dragonfly\n"),
	text.Colourf(utils.Broadcastprefix + "Subscribe to provsalt on <red>youtube</red>\n"),
	text.Colourf(utils.Broadcastprefix + "Check out our store at store.moonlightpe.com\n"),
	text.Colourf(utils.Broadcastprefix + "Rebirths are an essential part of the game /rebirth for more info!"),
	text.Colourf(utils.Broadcastprefix + "/help is a great place to start"),
}

func New() {
	for _, message := range Messages {
		time.Sleep(time.Minute * 8)
		_, _ = chat.Global.WriteString(message)
	}
}

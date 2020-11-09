package broadcast

import (
	"Prison/prisons/utils"
	"github.com/df-mc/dragonfly/dragonfly/player/chat"

	"github.com/sandertv/gophertunnel/minecraft/text"
	"time"
)

var Messages = [...]string{
	text.Colourf(utils.Broadcastprefix + "Join our discord today at discord.saltcraft.xyz\n"),
	text.Colourf(utils.Broadcastprefix + "Remember to vote for our server at vote.saltcraft.xyz for rewards\n"),
	text.Colourf(utils.Broadcastprefix + "Follow our twitter @saltcraft\n"),
	text.Colourf(utils.Broadcastprefix + "Check dragonfly out, our server software at github.com/df-mc/dragonfly\n"),
	text.Colourf(utils.Broadcastprefix + "Subscribe to provsalt on <red>youtube</red>\n"),
	text.Colourf(utils.Broadcastprefix + "Check out our store at store.saltcraft.xyz\n"),
	text.Colourf(utils.Broadcastprefix + "Rebirths are an essential part of the game /rebirth for more info!"),
	text.Colourf(utils.Broadcastprefix + "/help is a great place to start"),
}

func New() {
	for _, message := range Messages {
		time.Sleep(time.Minute * 8)
		chat.Global.Printf(message)
	}
}

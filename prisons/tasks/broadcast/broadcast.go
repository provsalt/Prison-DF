package broadcast

import (
	"github.com/df-mc/dragonfly/dragonfly/player/chat"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"time"
)

const broadcastprefix = "<grey>[<green><b!</b></green>]</grey> "

var Messages = [...]string{
	text.Colourf(broadcastprefix + "Join our discord today at discord.saltcraft.xyz\n"),
	text.Colourf(broadcastprefix + "Remember to vote for our server at vote.saltcraft.xyz for rewards\n"),
	text.Colourf(broadcastprefix + "Follow our twitter @saltcraft\n"),
	text.Colourf(broadcastprefix + "Check dragonfly out, our server software at github.com/df-mc/dragonfly\n"),
	text.Colourf(broadcastprefix + "Subscribe to provsalt on <red>youtube</red>\n"),
	text.Colourf(broadcastprefix + "Check out our store at store.saltcraft.xyz"),
}

func New() {
	for _, message := range Messages {
		time.Sleep(time.Minute * 6)
		chat.Global.Printf(message)
	}
	for _, message := range Messages {
		time.Sleep(time.Minute * 6)
		chat.Global.Printf(message)
	}
}

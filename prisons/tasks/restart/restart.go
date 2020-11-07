package restart

import (
	"Prison/prisons/utils"
	"github.com/df-mc/dragonfly/dragonfly/player/chat"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"time"
)

func New() {
	uptime := utils.Server.Uptime().Round(time.Minute).Minutes()

	switch utils.Server.Uptime().Round(time.Minute).Minutes() {
	case 30, 40, 50:
		chat.Global.Printf(text.Colourf(utils.Broadcastprefix+"Server will restart in %v minutes\n", 60-uptime))
	case 55, 56, 57, 58, 59:
		chat.Global.Printf(text.Colourf(utils.Broadcastprefix+"<red>Server will restart in %v minutes</red>\n", 60-uptime))
	case 60:
		chat.Global.Printf(text.Colourf(utils.Broadcastprefix + "<red>Server restarts now</red>\n"))
		time.Sleep(time.Second * 3)
		_ = utils.Server.Close()
	}
}

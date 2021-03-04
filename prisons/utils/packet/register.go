package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol/packet"

// Register registers all the custom packets used.
func Register() {
	packet.Register(IDPlayerInfo, func() packet.Packet {
		return &PlayerInfo{}
	})
}

package packet

import (
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const IDPlayerInfo = 0xfa

type PlayerInfo struct {
	UUID    uuid.UUID
	Address string
	XUID    string
}

func (p *PlayerInfo) ID() uint32 {
	return IDPlayerInfo
}

func (p *PlayerInfo) Marshal(w *protocol.Writer) {
	w.UUID(&p.UUID)
	w.String(&p.Address)
	w.String(&p.XUID)
}

func (p *PlayerInfo) Unmarshal(r *protocol.Reader) {
	r.UUID(&p.UUID)
	r.String(&p.Address)
	r.String(&p.XUID)
}

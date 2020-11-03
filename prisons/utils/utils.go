package utils

import (
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/session"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
	_ "unsafe"
)

var Server *dragonfly.Server

var Logger *logrus.Logger

func GetPrefix() string {
	return text.Colourf("<b><green>Salt</green><yellow>Craft</yellow></b>")
}

func GetServer() *dragonfly.Server {
	return Server
}

func GetLogger() *logrus.Logger {
	return Logger
}

func Vec64To32(vec3 mgl64.Vec3) mgl32.Vec3 {
	return mgl32.Vec3{float32(vec3[0]), float32(vec3[1]), float32(vec3[2])}
}

//go:linkname Player_session github.com/df-mc/dragonfly/dragonfly/player.(*Player).session
//noinspection ALL
func Player_session(*player.Player) *session.Session

//go:linkname Session_connection github.com/df-mc/dragonfly/dragonfly/session.(*Session).connection
//noinspection ALL
func Session_connection(*session.Session) *minecraft.Conn

//go:linkname Session_writePacket github.com/df-mc/dragonfly/dragonfly/session.(*Session).writePacket
//noinspection ALL
func Session_writePacket(*session.Session, packet.Packet)

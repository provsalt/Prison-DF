package utils

import (
	"Prison/economy"
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/session"
	worldmanager2 "github.com/emperials/df-worldmanager"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/sirupsen/logrus"
	_ "unsafe"
)

var Server *dragonfly.Server

var Logger *logrus.Logger

var Worldmanager *worldmanager2.WorldManager

var Economy *economy.Economy

func GetServer() *dragonfly.Server {
	return Server
}

func GetLogger() *logrus.Logger {
	return Logger
}

func Vec64To32(vec3 mgl64.Vec3) mgl32.Vec3 {
	return mgl32.Vec3{float32(vec3[0]), float32(vec3[1]), float32(vec3[2])}
}

func GetWorldmanager() *worldmanager2.WorldManager {
	return Worldmanager
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

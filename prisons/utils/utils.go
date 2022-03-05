package utils

import (
	"Prison/prisons/database/economy"
	"Prison/prisons/database/punishment"
	"Prison/prisons/database/ranks"
	"Prison/prisons/database/userinfo"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-plus/worldmanager"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/sirupsen/logrus"
	_ "unsafe"
)

var Server *server.Server

var Logger *logrus.Logger

var Worldmanager *worldmanager.WorldManager

var Economy *economy.Economy

var Ranks *ranks.RankApi

var (
	RanksDB     *ranks.RankApi
	EconomyDB   *economy.Economy
	Punishments *punishment.Database
	UserDB      *userinfo.UserInfo
)

// Development mode
var Development bool

const (
	WebhookURL = "https://discord.com/api/webhooks/791339145364111370/F1l9IYSkK3xDNhtf7Qc4tfhVqwIU0ACxUgMFU_QzOdfgPk6syKRYkWWT3k3ctydjY_JJ"
)

func GetServer() *server.Server {
	return Server
}

func GetLogger() *logrus.Logger {
	return Logger
}

func Vec64To32(vec3 mgl64.Vec3) mgl32.Vec3 {
	return mgl32.Vec3{float32(vec3[0]), float32(vec3[1]), float32(vec3[2])}
}

func GetWorldmanager() *worldmanager.WorldManager {
	return Worldmanager
}

//go:linkname Player_session github.com/df-mc/dragonfly/server/player.(*Player).session
//noinspection ALL
func Player_session(*player.Player) *session.Session

//go:linkname Session_connection github.com/df-mc/dragonfly/server/session.(*Session).connection
//noinspection ALL
func Session_connection(*session.Session) *minecraft.Conn

//go:linkname Session_writePacket github.com/df-mc/dragonfly/server/session.(*Session).writePacket
//noinspection ALL
func Session_writePacket(*session.Session, packet.Packet)

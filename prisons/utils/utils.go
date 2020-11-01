package utils

import (
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/session"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
	"regexp"
	_ "unsafe"
)

var Server *dragonfly.Server

var Logger *logrus.Logger

func GetPrefix() string {
	return text.Colourf("<b><green>Salt</green><yellow>Craft</yellow></b>")
}

func Colorize(message string) string {
	r := regexp.MustCompile("/&([0-9a-fk-or])/u")
	return r.ReplaceAllString(message, "§$1"+message)
}
func GetServer() *dragonfly.Server {
	return Server
}

func GetLogger() *logrus.Logger {
	return Logger
}

//go:linkname player_session github.com/df-mc/dragonfly/dragonfly/player.(*Player).session
//noinspection ALL
func Player_session(*player.Player) *session.Session

//go:linkname session_connection github.com/df-mc/dragonfly/dragonfly/session.(*Session).connection
//noinspection ALL
func Session_connection(*session.Session) *minecraft.Conn

//go:linkname session_writePacket github.com/df-mc/dragonfly/dragonfly/session.(*Session).writePacket
//noinspection ALL
func Session_writePacket(*session.Session, packet.Packet)

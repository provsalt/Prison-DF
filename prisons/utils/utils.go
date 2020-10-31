package utils

import (
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
	"regexp"
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

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
	return text.Bold()(text.Green()("Salt") + text.Yellow()("Craft"))
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

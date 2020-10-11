package utils

import (
	"github.com/sandertv/gophertunnel/minecraft/text"
	"regexp"
)

func GetPrefix() string {
	return text.Bold()(text.Green()("Salt") + text.Yellow()("Craft"))
}

func Colorize(message string) string {
	r := regexp.MustCompile("/&([0-9a-fk-or])/u")
	return r.ReplaceAllString(message, "§$1"+message)
}

package utils

import "github.com/sandertv/gophertunnel/minecraft/text"

const Broadcastprefix = "<grey>[<green><b>!</b></green>]</grey> "

func GetPrefix() string {
	return text.Colourf("<b><grey>Moon</grey><white>Light</white></b> ")
}

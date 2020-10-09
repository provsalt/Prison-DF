package commands

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Help struct {
	pages string
}

func (h Help) Run(source cmd.Source, output *cmd.Output) {
	output.Printf(text.White()("This is a test"))
}

package commands

import "github.com/df-mc/dragonfly/dragonfly/cmd"

type Gamemode struct{}

func (g Gamemode) Run(source cmd.Source, output *cmd.Output) {
	output.Print(source)
}

package punish

import (
	"Prison/prisons/console"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
)

type Punish struct {
	Sub UI
}

type UI string

func (u UI) SubName() string {
	return "ui"
}

func (p Punish) Run(source cmd.Source, output *cmd.Output) {
	if _, ok := source.(*console.Console); ok {
		output.Printf("UIs cannot be used by console")
	}
}

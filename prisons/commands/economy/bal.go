package economy

import (
	"Prison/prisons/utils"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Bal struct {
	Sub    bal
	Target []cmd.Target `optional:""`
}

type bal string

func (b Bal) Run(source cmd.Source, output *cmd.Output) {
	if len(b.Target) == 0 {
		p, ok := source.(*player.Player)
		if !ok {
			return
		}
		err, bal := utils.Economy.Balance(p)

		if err != nil {
			output.Printf(err.Error())
			return
		}
		output.Printf(text.Colourf("<yellow>Your balamce is %v</yellow> ", bal))
	}
	if len(b.Target) > 1 {
		output.Printf(text.Colourf("<red>You cannot run this command with this many people selected</red>"))
		return
	}
	for _, target := range b.Target {
		p, ok := target.(*player.Player)
		if !ok {
			output.Printf(text.Colourf("<red>This player does not exist</red>"))
		}
		err, bal := utils.Economy.Balance(p)

		if err != nil {
			output.Printf(err.Error())
			return
		}
		output.Printf(text.Colourf("<yellow>%s balamce is %v</yellow> ", target.Name(), bal))
	}
}

func (b bal) SubName() string {
	return "bal"
}

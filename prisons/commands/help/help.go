package help

import (
	"Prison/prisons/PlayerQuit/help"
	"Prison/prisons/utils"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/player/form"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Help struct{}

func (h Help) Run(source cmd.Source, output *cmd.Output) {
	if _, ok := source.(*player.Player); !ok {
		output.Printf(text.ANSI(utils.GetPrefix() + " How'd you forget the commands dumbass"))
		return
	}

	p := source.(*player.Player)
	f := form.NewMenu(help.HelpForm{
		CloseButton: form.Button{Text: "Okay"},
	}, "test")
	f = f.WithBody(
		text.Yellow()("==========") + text.Green()(" Help ") + text.Yellow()("==========\n") +
			text.Green()("/rankup &brank up to the next mine"+
				"\n"+
				"/mine [Mine A-Z] &bTeleport you to your mine"+
				"\n"+
				"/mymoney &bChecks your money"),
	)
	p.SendForm(f)
}

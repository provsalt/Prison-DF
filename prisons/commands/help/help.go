package help

import (
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
	f := form.NewMenu(HelpForm{
		CloseButton: form.Button{Text: "Okay"},
	}, "test")
	f = f.WithBody(
		text.Colourf("<yellow>==========</yellow><green>Help</green><yellow>==========</yellow>\n" +
			"<green>/rankup <aqua>brank up to the next mine</aqua>\n" +
			"/mine [Mine A-Z] <aqua>Teleport you to your mine</aqua>\n " +
			"/mymoney <aqua>Checks your money</aqua></green>"),
	)
	p.SendForm(f)
}

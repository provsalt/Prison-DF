package economy

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	form2 "github.com/df-mc/dragonfly/dragonfly/player/form"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Economy struct {
	Sub SubHelp
}

type EconomyHelp struct {
}

type SubHelp string

func (e Economy) Run(source cmd.Source, output *cmd.Output) {
	if p, ok := source.(*player.Player); !ok {
		form := form2.NewMenu(EconomyHelp{}, "Economy Help")
		form = form.WithBody(text.Colourf(""))
		p.SendForm(form)
	}
}

func (eh EconomyHelp) Submit(submitter form2.Submitter, pressed form2.Button) {

}

func (s SubHelp) SubName() string {
	return "help"
}

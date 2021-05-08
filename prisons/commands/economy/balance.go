package economy

import (
	"Prison/prisons/utils"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Balance struct {
	Sub    balance
	Player []cmd.Target `optional:""`
}

type balance string

func (b Balance) Run(source cmd.Source, output *cmd.Output) {
	if len(b.Player) == 0 {
		p, ok := source.(*player.Player)
		if !ok {
			return
		}
		err, bal := utils.EconomyDB.Balance(p)

		if err != nil {
			output.Printf(err.Error())
			return
		}
		output.Printf(text.Colourf("<yellow>Your balamce is %v</yellow> ", bal))
	}
	if len(b.Player) > 1 {
		output.Printf(text.Colourf("<red>You cannot run this command with this many people selected</red>"))
		return
	}
	for _, target := range b.Player {
		p, ok := target.(*player.Player)
		if !ok {
			output.Printf(text.Colourf("<red>This player does not exist</red>"))
		}
		err, bal := utils.EconomyDB.Balance(p)

		if err != nil {
			output.Printf(err.Error())
			return
		}
		output.Printf(text.Colourf("<yellow>%s balamce is %v</yellow> ", target.Name(), bal))
	}
}

func (b balance) SubName() string {
	return "balance"
}

package economy

import (
	"Prison/prisons/utils"
	"database/sql"
	"errors"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type SubPay string

type Pay struct {
	Sub    SubPay
	Player []cmd.Target
	Amount int
}

func (p Pay) Run(source cmd.Source, output *cmd.Output) {
	if len(p.Player) > 1 {
		output.Printf(text.Colourf("<red>You cannot run this command with this many people selected</red>"))
		return
	}
	if p.Amount < 1 {
		output.Printf(text.Colourf("<red>You should try and pay next time when you have money :'(</red>"))
	}
	for _, target := range p.Player {
		player2, ok := target.(*player.Player)
		if !ok {
			output.Printf(text.Colourf("<red>This player does not exist</red>"))
		}
		err, bal := utils.EconomyDB.Balance(player2)

		if err != nil {
			output.Printf(err.Error())
			return
		}

		if bal > p.Amount {
			err := utils.EconomyDB.ReduceMoney(player2, bal)
			if errors.Is(err, sql.ErrNoRows) {
				output.Printf(text.Colourf("<red>This player does not exist</red>"))
			} else {
				output.Printf(text.Colourf("<red>%s</red>", err.Error()))
			}
			err = utils.EconomyDB.AddMoney(player2, p.Amount)

			if errors.Is(err, sql.ErrNoRows) {
				output.Printf(text.Colourf("<red>This player does not exist</red>"))
			} else {
				output.Printf(text.Colourf("<red>%s</red>", err.Error()))
			}
		}
		output.Printf(text.Colourf("<green>Successfully sent %v to %s</green> ", p.Amount, target.Name()))
	}
}

func (s SubPay) SubName() string {
	return "pay"
}

package economy

import (
	"Prison/prisons/utils"
	"Prison/ranks"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type SubReduce string

type Reduce struct {
	Sub    SubReduce
	Player []cmd.Target
	Amount int
}

func (s SubReduce) SubName() string {
	return "reduce"
}

func (r Reduce) Run(source cmd.Source, output *cmd.Output) {
	if p, ok := source.(*player.Player); ok {
		if utils.Ranks.GetPermissionLevel(p).StaffRanks < ranks.Manager {
			output.Printf(text.Colourf("<red><You aren't allowed to use this command :(/red>"))
		}
	}

	if len(r.Player) > 1 {
		output.Printf(text.Colourf("<red>For security reasons, you are not allowed to give money to more then 1 player</red>"))
	}

	if r.Amount < 1 {
		output.Printf("<red>Invalid amount :(</red>")
	}

	for _, target := range r.Player {
		player2, ok := target.(*player.Player)
		if !ok {
			output.Printf(text.Colourf("<red>Sorry, I had trouble finding that player</red>"))
		}
		err := utils.Economy.ReduceMoney(player2, r.Amount)
		if err != nil {
			output.Printf("%v", err)
			utils.GetLogger().Errorln(err)
		}
	}
}

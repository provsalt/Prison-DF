package voting

import (
	"Prison/prisons/utils"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/player/title"
	"github.com/go-resty/resty/v2"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"time"
)

type Vote struct{}

func (v Vote) Run(source cmd.Source, output *cmd.Output) {
	if _, ok := source.(*player.Player); !ok {
		output.Printf(text.ANSI(utils.GetPrefix() + "You can't use this command stupid"))
		return
	}
	p := source.(*player.Player)
	err, res := CheckVote(p)
	if err != nil {
		utils.Logger.Errorln(err)
	}

	switch res {
	case "nil":
		output.Printf(text.Colourf(utils.GetPrefix() + "<red>You have not voted yet. To vote, head to vote.saltcraft.xyz</red>"))
	case "voted":
		Success(p)

		// Set vote as claimed
		r := resty.New()
		_, err := r.R().Get("https://minecraftpocket-servers.com/api/?action=post&object=votes&element=claim&key=FT01XbiS2IfonB16SKZ1jNKcNLID9fdEAk&username=" + p.Name())
		if err != nil {
			utils.GetLogger().Errorln(err)
		}
	case "claimed":
		output.Printf("You already have voted today")
	}
}

func CheckVote(player *player.Player) (error, string) {
	rest := resty.New()
	resp, err := rest.R().
		EnableTrace().
		Get("https://minecraftpocket-servers.com/api/?object=votes&element=claim&key=FT01XbiS2IfonB16SKZ1jNKcNLID9fdEAk&username=" + player.Name())
	if resp.Body() == nil {
		return nil, ""
	}
	switch string(resp.Body()) {
	case "0":
		return nil, "nil"
	case "1":
		return nil, "voted"
	case "2":
		return nil, "claimed"
	default:
		return err, ""
	}
}
func Success(player *player.Player) {
	t := title.New(text.Colourf("<green>Vote successful!</green>"))
	t = t.WithFadeOutDuration(time.Second * 3)
	t = t.WithSubtitle(text.Colourf("<aqua>Thank you for your support </aqua> <b><red><3</red></b>"))
	player.SendTitle(t)
}
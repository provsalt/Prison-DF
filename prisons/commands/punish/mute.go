package punish

import (
	"Prison/prisons/database/ranks"
	"Prison/prisons/utils"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xhit/go-str2duration/v2"
	"time"
)

type SubMute string

type MutePlayer struct {
	Sub    SubMute
	Target []cmd.Target
	Time   string
	Reason string `optional:""`
}

func (m SubMute) SubName() string {
	return "mute"
}

func (m MutePlayer) Run(source cmd.Source, output *cmd.Output) {
	if p, ok := source.(*player.Player); ok {
		if utils.Ranks.GetPermissionLevel(p).StaffRanks < ranks.Helper {
			output.Errorf(text.Colourf("You do not have permissions to run this command."))
		}
	}

	if len(m.Target) > 1 {
		output.Errorf("You can only specify one player")
		return
	}
	p, ok := m.Target[0].(*player.Player)
	if !ok {
		output.Errorf("User %s not found", m.Target[0].Name())
		return
	}

	if m.Reason == "" {
		m.Reason = "Unspecific"
	}

	if m.Time == "inf" {
		year := time.Now().Year()
		hrs := int64((2100 - year) * 8760)
		dur := int64(time.Hour) * hrs
		Mute(p, m.Reason, time.Duration(dur), source)
	} else {
		dur, err := str2duration.ParseDuration(m.Time)
		if err != nil {
			output.Errorf(err.Error())
		}
		Mute(p, m.Reason, dur, source)
	}
}

func Mute(player *player.Player, reason string, duration time.Duration, staff cmd.Source) {
	// spew.Dump(player.Name(), reason, duration.Hours(), staff.Name())
}

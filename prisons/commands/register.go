package commands

import (
	gamemode2 "Prison/prisons/commands/gamemode"
	"Prison/prisons/commands/help"
	"Prison/prisons/commands/stop"
	"Prison/prisons/commands/version"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
)

func Register() bool {
	cmd.Register(cmd.New("version", "Allows the user to view the version of the server", []string{"ver", "about"}, version.Version{}))
	cmd.Register(cmd.New("help", "Provides helpful infomation out thwre", nil, help.Help{}))
	cmd.Register(cmd.New("gamemode", "Set your own gamemode", []string{"gm"}, gamemode2.Gamemode{}))
	cmd.Register(cmd.New("stop", "Stops the server", nil, stop.Stop{}))
	return true
}

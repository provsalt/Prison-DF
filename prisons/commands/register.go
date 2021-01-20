package commands

import (
	"Prison/prisons/commands/economy"
	gamemode2 "Prison/prisons/commands/gamemode"
	"Prison/prisons/commands/help"
	"Prison/prisons/commands/say"
	"Prison/prisons/commands/stop"
	"Prison/prisons/commands/version"
	"Prison/prisons/commands/voting"
	"Prison/prisons/commands/world"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
)

// Register ...
func Register() bool {
	cmd.Register(cmd.New("version", "Allows the user to view the version of the server", []string{"ver", "about"}, version.Version{}))
	cmd.Register(cmd.New("help", "Provides helpful infomation about the server", []string{"?"}, help.Help{}))
	cmd.Register(cmd.New("gamemode", "Set your own gamemode", []string{"gm"}, gamemode2.Gamemode{}))
	cmd.Register(cmd.New("stop", "Stops the server", nil, stop.Stop{}))
	cmd.Register(cmd.New("vote", "Vote for the server", nil, voting.Vote{}))
	cmd.Register(cmd.New("say", "Broadcast your message", nil, say.Say{}))
	cmd.Register(cmd.New("world", "Manage worlds for staff only", nil, world.Teleport{}))
	cmd.Register(cmd.New("economy", "The economy commands", []string{"eco", "e"}, economy.Economy{}, economy.Bal{}, economy.Balance{}, economy.Pay{}, economy.Reduce{}))
	// cmd.Register()

	return true
}

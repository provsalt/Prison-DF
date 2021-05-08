package world

import (
	"Prison/prisons/utils"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Teleport struct {
	Sub   teleport
	World string
}

func (t Teleport) Run(source cmd.Source, output *cmd.Output) {
	p, ok := source.(*player.Player)
	if ok {
		w, ok := utils.GetWorldmanager().World(t.World)
		if !ok {
			output.Printf("This world does not exist")
		}
		w.AddEntity(p)
		p.Teleport(w.Spawn().Vec3Middle())
	}
}

type teleport string

func (teleport) SubName() string {
	return "teleport"
}

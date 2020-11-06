package console

// Eren5960 <ahmederen123@gmail.com>
import (
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Console struct{}

func (c Console) Name() string {
	return "Console"
}

func (c Console) Position() mgl64.Vec3 {
	return [3]float64{}
}

func (c Console) World() *world.World {
	return nil
}

func (Console) SendCommandOutput(output *cmd.Output) {
	for _, m := range output.Messages() {
		fmt.Println(m)
	}

	for _, e := range output.Errors() {
		fmt.Println(e.Error())
	}
}

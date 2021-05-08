package say

import (
	"Prison/prisons/console"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Say struct {
	Message cmd.Varargs
}

func (s Say) Run(source cmd.Source, output *cmd.Output) {
	_, ok := source.(*console.Console)

	if ok {
		if len(s.Message) == 0 {
			return
		}
		_, _ = chat.Global.WriteString(text.Colourf("<b><red>CONSOLE: </red></b>%s\n", s.Message))
	}
}

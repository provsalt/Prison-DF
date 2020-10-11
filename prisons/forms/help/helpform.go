package help

import (
	"github.com/df-mc/dragonfly/dragonfly/player/form"
)

type HelpForm struct {
	CloseButton form.Button
}

func (h HelpForm) Submit(submitter form.Submitter, pressed form.Button) {
	if pressed.Text == h.CloseButton.Text {
		return
	}
}

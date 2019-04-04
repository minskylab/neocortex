package terminal

import "github.com/bregydoc/neocortex"

type InputText struct {
	Text string
}

func (in *InputText) Type() neocortex.PrimitiveInputType {
	return neocortex.PrimitiveInputText
}

func (in *InputText) Value() string {
	return in.Text
}

func (in *InputText) Data() []byte {
	return []byte(in.Text)
}

func (term *Channel) NewInputText(c *neocortex.Context, text string, e []neocortex.Entity, i []neocortex.Intent) neocortex.Input {
	return &Input{
		inputType: &InputText{Text: text},
		c:         c,
		entities:  e,
		intents:   i,
	}
}

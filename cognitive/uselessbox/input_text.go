package uselessbox

import "github.com/bregydoc/neocortex"

type InputText struct {
	text string
}

func (in *InputText) Type() neocortex.PrimitiveInputType {
	return neocortex.PrimitiveInputText
}

func (in *InputText) Value() string {
	return in.text
}

func (in *InputText) Data() []byte {
	return []byte(in.text)
}

func (useless *Cognitive) NewInputText(c *neocortex.Context, text string) *Input {
	return &Input{
		context:   c,
		inputType: &InputText{text: text},
	}
}

package uselessbox

import neo "github.com/bregydoc/neocortex"

func (useless *Cognitive) NewInputText(c *neo.Context, text string, i []neo.Intent, e []neo.Entity) *neo.Input {
	t := neo.InputData{
		Type:  neo.InputText,
		Value: text,
		Data:  []byte(text),
	}
	return useless.NewInput(c, t, i, e)
}

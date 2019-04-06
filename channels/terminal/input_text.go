package terminal

import neo "github.com/bregydoc/neocortex"

func (term *Channel) NewInputText(c *neo.Context, text string, i []neo.Intent, e []neo.Entity) *neo.Input {
	t := neo.InputData{
		Type:  neo.InputText,
		Value: text,
		Data:  []byte(text),
	}
	return term.NewInput(c, t, i, e)
}

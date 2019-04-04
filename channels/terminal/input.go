package terminal

import neo "github.com/bregydoc/neocortex"

func (term *Channel) NewInput(c *neo.Context, inputType neo.InputType, i []neo.Intent, e []neo.Entity) *neo.Input {
	return &neo.Input{
		Context:   c,
		InputType: inputType,
		Intents:   i,
		Entities:  e,
	}
}

package facebook

import neo "github.com/bregydoc/neocortex"

func (fb *Channel) NewInput(c *neo.Context, inputType neo.InputType, i []neo.Intent, e []neo.Entity) *neo.Input {
	return &neo.Input{
		Context:   c,
		InputType: inputType,
		Intents:   i,
		Entities:  e,
	}
}

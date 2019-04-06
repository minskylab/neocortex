package terminal

import neo "github.com/bregydoc/neocortex"

func (term *Channel) NewInput(c *neo.Context, data neo.InputData, i []neo.Intent, e []neo.Entity) *neo.Input {
	return &neo.Input{
		Context:  c,
		Data:     data,
		Intents:  i,
		Entities: e,
	}
}

package uselessbox

import neo "github.com/bregydoc/neocortex"

func (useless *Cognitive) NewInput(c *neo.Context, data neo.InputData, i []neo.Intent, e []neo.Entity) *neo.Input {
	return &neo.Input{
		Context:  c,
		Data:     data,
		Intents:  i,
		Entities: e,
	}
}

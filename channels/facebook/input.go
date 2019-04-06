package facebook

import neo "github.com/bregydoc/neocortex"

func (fb *Channel) NewInput(c *neo.Context, data neo.InputData, i []neo.Intent, e []neo.Entity) *neo.Input {
	return &neo.Input{
		Context:  c,
		Data:     data,
		Intents:  i,
		Entities: e,
	}
}

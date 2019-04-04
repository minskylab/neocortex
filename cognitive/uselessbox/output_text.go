package uselessbox

import neo "github.com/bregydoc/neocortex"

func (useless *Cognitive) NewOutputText(c *neo.Context) *neo.Output {
	res := []neo.Response{{
		Type:     neo.Text,
		Value:    "I'm useless, you don't wait more from me",
		IsTyping: false,
	}}
	return useless.NewOutput(c, res, nil, nil)
}

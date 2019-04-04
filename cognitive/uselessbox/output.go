package uselessbox

import neo "github.com/bregydoc/neocortex"

func (useless *Cognitive) NewOutput(c *neo.Context, res []neo.Response, i []neo.Intent, e []neo.Entity) *neo.Output {
	return &neo.Output{
		Context:      c,
		Entities:     e,
		Intents:      i,
		Responses:    res,
		VisitedNodes: nil, // Proper of this Uselessbox
		Logs:         nil, // Proper of this Uselessbox
	}
}

package uselessbox

import "github.com/bregydoc/neocortex"

type Output struct {
	context *neocortex.Context
}

func NewOutput(c *neocortex.Context) *Output {
	return &Output{context: c}
}

func (out *Output) Context() *neocortex.Context {
	return out.context
}

func (out *Output) Entities() []neocortex.Entity {
	return []neocortex.Entity{}
}

func (out *Output) Intents() []neocortex.Intent {
	return []neocortex.Intent{}
}

func (out *Output) VisitedNodes() []*neocortex.DialogNode {
	return []*neocortex.DialogNode{}
}

func (out *Output) Logs() []*neocortex.LogMessage {
	return []*neocortex.LogMessage{}
}

func (out *Output) Responses() []neocortex.Response {
	return []neocortex.Response{
		NewOutputText(),
	}
}

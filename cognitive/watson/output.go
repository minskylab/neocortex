package watson

import (
	"github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

type Output struct {
	context *neocortex.Context
	r       *assistantv2.MessageResponse
}

func (watson *Cognitive) NewOutput(c *neocortex.Context, msn *assistantv2.MessageResponse) *Output {
	var ok bool
	c.Variables, ok = msn.Output.UserDefined.(map[string]interface{}) // TODO: Evaluate
	if !ok {
		c.Variables = map[string]interface{}{}
	}
	return &Output{
		context: c,
		r:       msn,
	}
}

func (out *Output) Context() *neocortex.Context {
	return out.context
}

func (out *Output) Entities() []neocortex.Entity {
	entities := make([]neocortex.Entity, 0)
	for _, e := range out.r.Output.Entities {
		entities = append(entities, getNeocortexEntity(e))
	}
	return entities
}

func (out *Output) Intents() []neocortex.Intent {
	intents := make([]neocortex.Intent, 0)
	for _, i := range out.r.Output.Intents {
		intents = append(intents, getNeocortexIntent(i))
	}
	return intents
}

func (out *Output) VisitedNodes() []*neocortex.DialogNode {
	nodes := make([]*neocortex.DialogNode, 0)
	for _, n := range out.r.Output.Debug.NodesVisited {
		nodes = append(nodes, &neocortex.DialogNode{
			Title:      *n.Title,
			Conditions: *n.Conditions,
			Name:       *n.DialogNode,
		})
	}
	return nodes
}

func (out *Output) Logs() []*neocortex.LogMessage {
	logs := make([]*neocortex.LogMessage, 0)
	for _, l := range out.r.Output.Debug.LogMessages {
		logs = append(logs,
			&neocortex.LogMessage{
				Level:   neocortex.LogLevelType(*l.Message),
				Message: *l.Level,
			})
	}
	return logs
}

func (out *Output) Responses() []neocortex.Response {
	responses := make([]neocortex.Response, 0)
	for _, gen := range out.r.Output.Generic {
		switch *gen.ResponseType {
		case string(neocortex.Text):
			responses = append(responses, out.NewOutputText(gen))
		default:
			responses = append(responses, out.NewOutputUnknown(gen))
		}
	}

	return responses
}

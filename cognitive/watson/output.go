package watson

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

func (watson *Cognitive) NewOutput(c *neo.Context, r *assistantv2.MessageResponse) *neo.Output {
	entities := make([]neo.Entity, 0)
	for _, e := range r.Output.Entities {
		entities = append(entities, getNeocortexEntity(e))
	}

	intents := make([]neo.Intent, 0)
	for _, i := range r.Output.Intents {
		intents = append(intents, getNeocortexIntent(i))
	}

	logs := make([]*neo.LogMessage, 0)
	for _, l := range r.Output.Debug.LogMessages {
		logs = append(logs,
			&neo.LogMessage{
				Level:   neo.LogLevelType(*l.Message),
				Message: *l.Level,
			})
	}

	nodes := make([]*neo.DialogNode, 0)
	for _, n := range r.Output.Debug.NodesVisited {
		nodes = append(nodes, &neo.DialogNode{
			Title:      *n.Title,
			Conditions: *n.Conditions,
			Name:       *n.DialogNode,
		})
	}

	responses := make([]neo.Response, 0)
	for _, gen := range r.Output.Generic {
		switch *gen.ResponseType {
		case "text":
			rText := watson.newTextResponse(gen)
			responses = append(responses, rText)
		case "option":
			rOption := watson.newOptionResponse(gen)
			responses = append(responses, rOption)
		default:
			rUnknown := watson.newUnknownResponse(gen)
			responses = append(responses, rUnknown)
		}
	}
	return &neo.Output{
		Context:      c,
		Logs:         logs,
		VisitedNodes: nodes,
		Intents:      intents,
		Entities:     entities,
		Responses:    responses,
	}
}

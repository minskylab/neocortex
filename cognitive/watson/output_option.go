package watson

import (
	"github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

func (watson *Cognitive) newOptionResponse(gen assistantv2.DialogRuntimeResponseGeneric) neocortex.Response {
	typing := false
	if gen.Typing != nil {
		typing = *gen.Typing
	}

	options := make([]*neocortex.Option, 0)

	for _, o := range gen.Options {
		options = append(options, &neocortex.Option{
			Text:   *o.Label,
			Action: *o.Value.Input.Text,
		})
	}
	return neocortex.Response{
		Type: neocortex.Options,
		Value: neocortex.OptionsResponse{
			Title:       *gen.Title,
			Description: *gen.Description,
			Options:     options,
		},
		IsTyping: typing,
	}
}

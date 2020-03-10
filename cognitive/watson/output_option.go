package watson

import (
	"strings"

	"github.com/minskylab/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

func (watson *Cognitive) newOptionResponse(gen assistantv2.DialogRuntimeResponseGeneric) neocortex.Response {
	typing := false
	if gen.Typing != nil {
		typing = *gen.Typing
	}

	options := make([]*neocortex.Option, 0)

	for _, o := range gen.Options {
		postBack := true
		if strings.HasPrefix(*o.Value.Input.Text, "http") {
			postBack = false
		}

		options = append(options, &neocortex.Option{
			Text:       *o.Label,
			Action:     *o.Value.Input.Text,
			IsPostBack: postBack,
		})
	}

	title := ""
	description := ""

	if gen.Title != nil {
		title = *gen.Title
	}

	if gen.Description != nil {
		description = *gen.Description
	}

	return neocortex.Response{
		Type: neocortex.Options,
		Value: neocortex.OptionsResponse{
			Title:       title,
			Description: description,
			Options:     options,
		},
		IsTyping: typing,
	}
}

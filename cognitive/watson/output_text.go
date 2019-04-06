package watson

import (
	"github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

func (watson *Cognitive) newTextResponse(gen assistantv2.DialogRuntimeResponseGeneric) neocortex.Response {
	typing := false
	if gen.Typing != nil {
		typing = *gen.Typing
	}
	return neocortex.Response{
		Type:     neocortex.Text,
		Value:    *gen.Text,
		IsTyping: typing,
	}
}

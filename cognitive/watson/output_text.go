package watson

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

func (watson *Cognitive) newTextResponse(gen assistantv2.RuntimeResponseGeneric) neo.Response {
	typing := false
	if gen.Typing != nil {
		typing = *gen.Typing
	}
	return neo.Response{
		Type:     neo.Text,
		Value:    *gen.Text,
		IsTyping: typing,
	}
}

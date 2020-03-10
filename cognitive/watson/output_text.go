package watson

import (
	neo "github.com/minskylab/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

func (watson *Cognitive) newTextResponse(gen assistantv2.DialogRuntimeResponseGeneric) neo.Response {
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

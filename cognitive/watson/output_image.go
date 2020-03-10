package watson

import (
	neo "github.com/minskylab/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

func (watson *Cognitive) newImageResponse(gen assistantv2.DialogRuntimeResponseGeneric) neo.Response {
	typing := false
	if gen.Typing != nil {
		typing = *gen.Typing
	}
	src := ""
	if gen.Source != nil {
		src = *gen.Source
	}
	return neo.Response{
		Type:     neo.Image,
		Value:    src,
		IsTyping: typing,
	}
}

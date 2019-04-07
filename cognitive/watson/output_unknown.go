package watson

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

func (watson *Cognitive) newUnknownResponse(gen assistantv2.DialogRuntimeResponseGeneric) neo.Response {
	return neo.Response{
		Type:     neo.Unknown,
		Value:    "unknown or not implemented type: " + *gen.ResponseType,
		IsTyping: false,
	}
}

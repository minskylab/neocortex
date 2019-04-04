package watson

import (
	"github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

func (watson *Cognitive) NewUnknownResponse(gen assistantv2.DialogRuntimeResponseGeneric) neocortex.Response {
	return neocortex.Response{
		Type:     neocortex.Unknown,
		Value:    "unknown or not implemented type: " + *gen.ResponseType,
		IsTyping: false,
	}
}

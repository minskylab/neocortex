package watson

import (
	"github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

type OutputUnknown struct {
	context *neocortex.Context
	generic assistantv2.DialogRuntimeResponseGeneric
}

func (out *Output) NewOutputUnknown(generic assistantv2.DialogRuntimeResponseGeneric) *OutputUnknown {
	return &OutputUnknown{
		context: out.context,
		generic: generic,
	}
}

func (out *OutputUnknown) Type() neocortex.ResponseType {
	return neocortex.Unknown
}

func (out *OutputUnknown) Value() interface{} {
	return "unknown or not implemented type: " + *out.generic.ResponseType
}

func (out *OutputUnknown) IsTyping() bool {
	return false
}

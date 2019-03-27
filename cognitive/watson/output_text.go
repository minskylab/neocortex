package watson

import (
	"github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

type OutputText struct {
	context *neocortex.Context
	generic assistantv2.DialogRuntimeResponseGeneric
}

func (out *Output) NewOutputText(generic assistantv2.DialogRuntimeResponseGeneric) *OutputText {
	return &OutputText{
		context: out.context,
		generic: generic,
	}
}

func (out *OutputText) Type() neocortex.ResponseType {
	return neocortex.Text
}

func (out *OutputText) Value() interface{} {
	return *out.generic.Text
}

func (out *OutputText) IsTyping() bool {
	return *out.generic.Typing
}

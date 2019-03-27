package uselessbox

import "github.com/bregydoc/neocortex"

type OutputText struct {
	response string
}

func NewOutputText(text ...string) *OutputText {
	return &OutputText{
		response: "I'm useless, you don't wait more from me",
	}
}

func (out *OutputText) IsTyping() bool {
	return false
}

func (out *OutputText) Type() neocortex.ResponseType {
	return neocortex.Text
}

func (out *OutputText) Value() interface{} {
	return out.response
}

package neocortex

type GenericOutputText struct {
	c        *Context
	entities []Entity
	intents  []Intent
	text     string
}

type OutputTextResponse struct {
	text string
}

func CreateNewGenericOutputText(c *Context, text string, intents []Intent, entities []Entity) *GenericOutputText {
	return &GenericOutputText{
		text:     text,
		c:        c,
		intents:  intents,
		entities: entities,
	}
}

func (r *OutputTextResponse) IsTyping() bool {
	return false
}

func (r *OutputTextResponse) Type() ResponseType {
	return Text
}

func (r *OutputTextResponse) Value() interface{} {
	return r.text
}

func (out *GenericOutputText) Context() *Context {
	return out.c
}

func (out *GenericOutputText) Entities() []Entity {
	return nil
}

func (out *GenericOutputText) Intents() []Intent {
	return nil
}

func (out *GenericOutputText) VisitedNodes() []*DialogNode {
	return nil
}

func (out *GenericOutputText) Logs() []*LogMessage {
	return nil
}

func (out *GenericOutputText) Responses() []Response {
	return []Response{&OutputTextResponse{
		text: out.text,
	}}
}

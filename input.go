package neocortex

type PrimitiveInputType string

var PrimitiveInputText PrimitiveInputType = "text"

type InputType interface {
	Type() PrimitiveInputType
	Value() string
	Data() []byte
}

// Input represent an Input for the cognitive service
type Input interface {
	Context() *Context
	InputType() InputType
	Entities() []Entity
	Intents() []Intent
}

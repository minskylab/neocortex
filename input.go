package neocortex

type PrimitiveInputType string

var PrimitiveInputText PrimitiveInputType = "text"

type InputType struct {
	Type  PrimitiveInputType
	Value string
	Data  []byte
}

// Input represent an Input for the cognitive service
type Input struct {
	Context   *Context
	InputType InputType
	Entities  []Entity
	Intents   []Intent
}

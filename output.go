package neocortex

// ResponseType define the types of generic response
type ResponseType string

// Text is a kind of generic response
var Text ResponseType = "text"

// Pause is a kind of generic response
var Pause ResponseType = "pause"

// Image is a kind of generic response
var Image ResponseType = "image"

// Option is a kind of generic response
var Option ResponseType = "option"

// ConnectToAgent is a kind of generic response
var ConnectToAgent ResponseType = "connect_to_agent"

var Suggestion ResponseType = "suggestion"

var Unknown ResponseType = "unknown"

type Response interface {
	IsTyping() bool
	Type() ResponseType
	Value() interface{} // TODO: Evaluate
}

// Output represents the response of an input from the cognitive service
type Output interface {
	Context() *Context
	Entities() []Entity
	Intents() []Intent
	VisitedNodes() []*DialogNode
	Logs() []*LogMessage
	Responses() []Response
}

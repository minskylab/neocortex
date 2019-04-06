package neocortex

// ResponseType define the types of generic response
type ResponseType string

// Text is a kind of generic response
const Text ResponseType = "text"

// Pause is a kind of generic response
const Pause ResponseType = "pause"

// Image is a kind of generic response
const Image ResponseType = "image"

// Options is a kind of generic response
const Options ResponseType = "option"

// ConnectToAgent is a kind of generic response
// var ConnectToAgent ResponseType = "connect_to_agent"

const Suggestion ResponseType = "suggestion"

const Unknown ResponseType = "unknown"

type Response struct {
	IsTyping bool
	Type     ResponseType
	Value    interface{} // TODO: Evaluate
}

// Output represents the response of an input from the cognitive service
type Output struct {
	Context      *Context
	Entities     []Entity
	Intents      []Intent
	VisitedNodes []*DialogNode
	Logs         []*LogMessage
	Responses    []Response
}

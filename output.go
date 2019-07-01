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
	IsTyping bool         `json:"is_typing"`
	Type     ResponseType `json:"type"`
	Value    interface{}  `json:"value"`
}

// Output represents the response of an input from the cognitive service
type Output struct {
	Entities     []Entity      `json:"entities"`
	Intents      []Intent      `json:"intents"`
	VisitedNodes []*DialogNode `json:"visited_nodes"`
	Logs         []*LogMessage `json:"logs"`
	Responses    []Response    `json:"responses"`
}

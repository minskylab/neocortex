package neocortex

import "time"

// ResponseGenericType define the types of generic response
type ResponseGenericType string

// Text is a kind of generic response
var Text ResponseGenericType = "text"

// Pause is a kind of generic response
var Pause ResponseGenericType = "pause"

// Image is a kind of generic response
var Image ResponseGenericType = "image"

// Option is a kind of generic response
var Option ResponseGenericType = "option"

// ConnectToAgent is a kind of generic response
var ConnectToAgent ResponseGenericType = "connect_to_agent"

// var Suggestion ResponseGenericType = "suggestion"

// ResponseGeneric wraps and define a response from cognitive service
type ResponseGeneric struct {
	ResponseType        ResponseGenericType
	Text                string
	Time                time.Time
	Typing              bool
	Source              string
	Description         string
	Preference          string
	Options             map[string]string
	MessageToHumanAgent string
	Topic               string
	// Suggestion is a future feature
}

// Output represents the response of an input from the cognitive service
type Output struct {
	InputText    string
	Context      Context
	Entities     []*Entity
	Intents      []*Intent
	NodesVisited []*DialogNode
	// Actions         []*Action
	Actions  map[string]string
	Logs     []*LogMessage
	Response []*ResponseGeneric
}

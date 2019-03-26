package neocortex

// Input represent an Input for the cognitive service
type Input struct {
	Text         string
	Context      Context
	Entities     []Entity
	Intents      []Intent
	NodesVisited []DialogNode
}

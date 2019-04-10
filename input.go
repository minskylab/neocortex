package neocortex

type InputType string

const InputText InputType = "text"
const InputAudio InputType = "audio"
const InputImage InputType = "image"
const InputEmoji InputType = "emoji"

type InputData struct {
	Type  InputType `json:"type"`
	Value string    `json:"value"`
	Data  []byte    `json:"data"`
}

// Input represent an Input for the cognitive service
type Input struct {
	Context  *Context  `json:"context"`
	Data     InputData `json:"data"`
	Entities []Entity  `json:"entities"`
	Intents  []Intent  `json:"intents"`
}

package neocortex

type InputType string

const InputText InputType = "text"
const InputAudio InputType = "audio"
const InputImage InputType = "image"
const InputEmoji InputType = "emoji"

type InputData struct {
	Type  InputType
	Value string
	Data  []byte
}

// Input represent an Input for the cognitive service
type Input struct {
	Context  *Context
	Data     InputData
	Entities []Entity
	Intents  []Intent
}

package neocortex

type CognitiveService interface {
	GetProtoResponse(in *Input) (*Output, error)
}

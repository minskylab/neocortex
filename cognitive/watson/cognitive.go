package watson

import (
	"context"
	"github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

type Cognitive struct {
	service     *assistantv2.AssistantV2
	assistantID string
}

type NewCognitiveParams struct {
	Url         string
	Username    string
	Password    string
	Version     string
	AssistantID string
	ApiKey      string
}

func NewCognitive(params NewCognitiveParams) (*Cognitive, error) {
	assistant, err := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
		Version:   params.Version,
		Username:  params.Username,
		Password:  params.Password,
		URL:       params.Url,
		IAMApiKey: params.ApiKey,
	})
	if err != nil {
		return nil, err
	}
	return &Cognitive{
		service:     assistant,
		assistantID: params.AssistantID,
	}, nil
}

func (watson *Cognitive) CreateNewContext(c *context.Context, info neocortex.PersonInfo) *neocortex.Context {
	r, responseErr := watson.service.CreateSession(watson.service.NewCreateSessionOptions(watson.assistantID))
	if responseErr != nil {
		panic(responseErr)
	}
	sess := watson.service.GetCreateSessionResult(r)
	return &neocortex.Context{
		SessionID: *sess.SessionID,
		Person:    info,
		Context:   c,
		Variables: map[string]interface{}{},
	}
}

func (watson *Cognitive) GetProtoResponse(in *neocortex.Input) (*neocortex.Output, error) {
	var opts *assistantv2.MessageOptions
	switch in.Data.Type {
	// Watson only supports one type of input: InputText
	case neocortex.InputText:
		_, opts = watson.NewInputText(in.Context, in.Data.Value, in.Intents, in.Entities)
	default:
		return nil, neocortex.ErrInvalidInputType
	}

	r, err := watson.service.Message(opts)
	if err != nil {
		return nil, neocortex.ErrSessionNotExist
	}

	if r.StatusCode != 200 {
		return nil, neocortex.ErrInvalidResponseFromCognitiveService
	}

	response := watson.service.GetMessageResult(r)
	// pp.Println(response)
	out := watson.NewOutput(in.Context, response)

	return out, nil

}

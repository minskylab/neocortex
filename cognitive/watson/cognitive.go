package watson

import (
	"context"

	neo "github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

type Cognitive struct {
	service              *assistantv2.AssistantV2
	assistantID          string
	doneContextCallbacks []*func(c *neo.Context)
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
		// turnsMap:    map[string]int{},
	}, nil
}

func (watson *Cognitive) CreateNewContext(c *context.Context, info neo.PersonInfo) *neo.Context {
	r, responseErr := watson.service.CreateSession(watson.service.NewCreateSessionOptions(watson.assistantID))
	if responseErr != nil {
		panic(responseErr)
	}
	sess := watson.service.GetCreateSessionResult(r)

	// watson.turnsMap[*sess.SessionID] = 1
	return &neo.Context{
		SessionID: *sess.SessionID,
		Person:    info,
		Context:   c,
		Variables: map[string]interface{}{},
	}
}

func (watson *Cognitive) GetProtoResponse(c *neo.Context, in *neo.Input) (*neo.Output, error) {

	var opts *assistantv2.MessageOptions
	switch in.Data.Type {

	// Watson only supports one type of input: InputText
	case neo.InputText:
		_, opts = watson.NewInputText(in.Data.Value, c, in.Intents, in.Entities)
	default:
		return nil, neo.ErrInvalidInputType
	}

	r, err := watson.service.Message(opts)
	if err != nil {
		for _, call := range watson.doneContextCallbacks {
			(*call)(c)
		}
		return nil, neo.ErrSessionNotExist
	}

	if r.StatusCode != 200 {
		for _, call := range watson.doneContextCallbacks {
			(*call)(c)
		}
		return nil, neo.ErrInvalidResponseFromCognitiveService
	}

	response := watson.service.GetMessageResult(r)

	out := watson.NewOutput(c, response)

	return out, nil

}

func (watson *Cognitive) OnContextIsDone(callback func(c *neo.Context)) {
	if watson.doneContextCallbacks == nil {
		watson.doneContextCallbacks = []*func(c *neo.Context){}
	}
	watson.doneContextCallbacks = append(watson.doneContextCallbacks, &callback)
}

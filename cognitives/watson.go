package cognitives

import (
	"context"
	"errors"
	"github.com/jeremywohl/flatten"
	"github.com/watson-developer-cloud/go-sdk/core"

	"github.com/watson-developer-cloud/go-sdk/assistantv2"

	"github.com/bregydoc/neocortex"

	"time"
)

// WatsonCognitiveService represents a Watson Assistant cognitive service
type WatsonCognitiveService struct {
	AssistantID string
	assistant   *assistantv2.AssistantV2
	// session *assistantv2.SessionResponse
}

// CreateNewWatsonCognitive create a new watson based cognitive service
func CreateNewWatsonCognitive(url, username, password, version, assistantID string) (*WatsonCognitiveService, error) {
	assistant, err := assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
		Version:  version,
		Username: username,
		Password: password,
		URL:      url,
	})
	if err != nil {
		return nil, err
	}

	return &WatsonCognitiveService{
		assistant:   assistant,
		AssistantID: assistantID,
	}, nil
}

func (watson *WatsonCognitiveService) CreateNewSession(c context.Context, userID string) *neocortex.Context {
	response, responseErr := watson.assistant.CreateSession(watson.assistant.NewCreateSessionOptions(watson.AssistantID))
	if responseErr != nil {
		panic(responseErr)
	}
	sess := watson.assistant.GetCreateSessionResult(response)
	return &neocortex.Context{
		SessionID: *sess.SessionID,
		Context:   c,
		Metadata:  map[string]string{},
	}
}

func (watson *WatsonCognitiveService) getNativeEntity(e *neocortex.Entity) *assistantv2.RuntimeEntity {
	ent := &assistantv2.RuntimeEntity{
		Confidence: &e.Confidence,
		Entity:     &e.Entity,
		Location:   e.Location,
		Value:      &e.Value,
		Metadata:   e.Metadata,
	}
	// TODO: Groups work in progress
	return ent
}

func (watson *WatsonCognitiveService) getNativeIntent(i *neocortex.Intent) *assistantv2.RuntimeIntent {
	intent := &assistantv2.RuntimeIntent{
		Confidence: &i.Confidence,
		Intent:     &i.Intent,
	}

	return intent
}

func (watson *WatsonCognitiveService) getNativeIn(c *neocortex.Context, in *neocortex.Input) *assistantv2.MessageOptions {
	input := &assistantv2.MessageInput{
		Text: &in.Text,
	}

	entities := make([]assistantv2.RuntimeEntity, 0)
	for _, e := range in.Entities {
		entities = append(entities, *watson.getNativeEntity(e))
	}
	intents := make([]assistantv2.RuntimeIntent, 0)
	for _, i := range in.Intents {
		intents = append(intents, *watson.getNativeIntent(i))

	}

	opts := watson.assistant.NewMessageOptions(watson.AssistantID, c.SessionID)
	opts.SetInput(input)

	return opts
}

func (watson *WatsonCognitiveService) getNeoCortexOut(c *neocortex.Context, res *assistantv2.MessageResponse) *neocortex.Output {

	intents := make([]*neocortex.Intent, 0)
	for _, i := range res.Output.Intents {
		intents = append(intents, &neocortex.Intent{
			Intent:     *i.Intent,
			Confidence: *i.Confidence,
		})
	}

	entities := make([]*neocortex.Entity, 0)
	for _, e := range res.Output.Entities {
		metadata, ok := e.Metadata.(map[string]string)
		if !ok {
			metadata = make(map[string]string)
		}
		entities = append(entities, &neocortex.Entity{
			Entity:     *e.Entity,
			Confidence: *e.Confidence,
			Value:      *e.Value,
			Location:   e.Location,
			Metadata:   metadata,
			// Groups work in progress
		})
	}

	genResponses := make([]*neocortex.ResponseGeneric, 0)
	for _, g := range res.Output.Generic {
		if g.Text == nil {
			g.Text = core.StringPtr("")
		}

		if g.Typing == nil {
			g.Typing = core.BoolPtr(false)
		}

		if g.Source == nil {
			g.Source = core.StringPtr("")
		}

		if g.Time == nil {
			g.Time = core.Int64Ptr(0)
		}

		if g.Topic == nil {
			g.Topic = core.StringPtr("")
		}

		if g.Description == nil {
			g.Description = core.StringPtr("")
		}

		if g.ResponseType == nil {
			g.ResponseType = core.StringPtr("")
		}

		if g.MessageToHumanAgent == nil {
			g.MessageToHumanAgent = core.StringPtr("")
		}

		if g.Preference == nil {
			g.Preference = core.StringPtr("")
		}
		genResponses = append(genResponses, &neocortex.ResponseGeneric{
			Text:                *g.Text,
			Typing:              *g.Typing,
			Source:              *g.Source,
			Time:                time.Unix(*g.Time, 0),
			Topic:               *g.Topic,
			Description:         *g.Description,
			ResponseType:        neocortex.ResponseGenericType(*g.ResponseType),
			MessageToHumanAgent: *g.MessageToHumanAgent,
			Preference:          *g.Preference,
		})
	}

	var pseudoActions map[string]string
	if res.Output.UserDefined != nil {
		pseudoActions = make(map[string]string)
		definitions, err := flatten.Flatten(res.Output.UserDefined.(map[string]interface{}), "", flatten.DotStyle)
		if err != nil {
			panic(err)
		}
		for k, d := range definitions {
			info, ok := d.(string)
			if !ok {
				info = ""
			}
			pseudoActions[k] = info
		}
	}

	return &neocortex.Output{
		Context:  *c,
		Actions:  pseudoActions,
		Intents:  intents,
		Entities: entities,
		Response: genResponses,
	}
}

func (watson *WatsonCognitiveService) GetProtoResponse(c *neocortex.Context, in *neocortex.Input) (*neocortex.Output, error) {
	opts := watson.getNativeIn(c, in)
	// pp.Println(opts)
	r, err := watson.assistant.Message(opts)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, errors.New("err0xF3")
	}

	response := watson.assistant.GetMessageResult(r)

	out := watson.getNeoCortexOut(c, response)
	return out, nil

}

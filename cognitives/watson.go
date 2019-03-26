package cognitives

import (
	"errors"
	"github.com/k0kubun/pp"

	"github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv1"
	"time"
)

// WatsonCognitiveService represents a Watson Assistant cognitive service
type WatsonCognitiveService struct {
	workspaceID string
	assistant   *assistantv1.AssistantV1
}

// CreateNewWatsonCognitive create a new watson based cognitive service
func CreateNewWatsonCognitive(url, username, password, version, workspace string) (*WatsonCognitiveService, error) {
	assistant, err := assistantv1.NewAssistantV1(&assistantv1.AssistantV1Options{
		Version:  version,
		Username: username,
		Password: password,
		URL:      url,
	})
	if err != nil {
		return nil, err
	}

	// r, err := assistant.CreateWorkspace(
	// 	&assistantv1.CreateWorkspaceOptions{
	// 		Name: core.StringPtr("API test"),
	// 		Description: core.StringPtr("Example workspace created via API"),
	// 	},
	// )
	// if err != nil {
	// 	return nil, err
	// }
	//
	// ws := assistant.GetCreateWorkspaceResult(r)
	// pp.Println(ws)
	return &WatsonCognitiveService{
		workspaceID: workspace,
		assistant:   assistant,
	}, nil
}

func (watson *WatsonCognitiveService) getNativeEntity(e *neocortex.Entity) *assistantv1.RuntimeEntity {
	ent := &assistantv1.RuntimeEntity{}
	// ent.SetMetadata(&(e.Metadata))
	ent.SetConfidence(&e.Confidence)
	ent.SetEntity(&e.Entity)
	ent.SetLocation(&e.Location)
	ent.SetValue(&e.Value)
	// Groups work in progress
	return ent
}

func (watson *WatsonCognitiveService) getNativeIntent(i *neocortex.Intent) *assistantv1.RuntimeIntent {
	intent := &assistantv1.RuntimeIntent{}
	intent.SetConfidence(&i.Confidence)
	intent.SetIntent(&i.Intent)
	return intent
}

func (watson *WatsonCognitiveService) getNativeIn(in *neocortex.Input) *assistantv1.MessageOptions {
	input := &assistantv1.InputData{}
	input.SetText(&in.Text)
	context := &assistantv1.Context{}
	context.SetConversationID(&in.Context.ConversationID)
	entities := make([]assistantv1.RuntimeEntity, 0)
	for _, e := range in.Entities {
		entities = append(entities, *watson.getNativeEntity(e))
	}
	intents := make([]assistantv1.RuntimeIntent, 0)
	for _, i := range in.Intents {
		intents = append(intents, *watson.getNativeIntent(i))

	}
	output := &assistantv1.OutputData{}
	nodes := make([]string, 0)
	nodesDetails := make([]assistantv1.DialogNodeVisitedDetails, 0)
	for _, n := range in.NodesVisited {
		nodes = append(nodes, n.Name)
		l := assistantv1.DialogNodeVisitedDetails{
			Title:      &n.Title,
			Conditions: &n.Conditions,
			DialogNode: &n.Name,
		}
		nodesDetails = append(nodesDetails, l)
	}
	output.SetNodesVisited(&nodes)
	output.SetNodesVisitedDetails(&nodesDetails)

	return &assistantv1.MessageOptions{
		WorkspaceID: &watson.workspaceID,
		Context:     context,
		Input:       input,
		Entities:    entities,
		Intents:     intents,
		Output:      output,
	}
}

func (watson *WatsonCognitiveService) getNeoCortexOut(res *assistantv1.MessageResponse) *neocortex.Output {

	// conversationID := *(res.Context)["conversation_id"]
	context := neocortex.Context{
		// ConversationID: *conversationID,
	}

	intents := make([]*neocortex.Intent, 0)
	for _, i := range res.Intents {
		intents = append(intents, &neocortex.Intent{
			Intent:     *i.GetIntent(),
			Confidence: *i.GetConfidence(),
		})
	}

	entities := make([]*neocortex.Entity, 0)
	for _, e := range res.Entities {
		entities = append(entities, &neocortex.Entity{
			Entity:     *e.GetEntity(),
			Confidence: *e.GetConfidence(),
			Value:      *e.GetValue(),
			Location:   *e.GetLocation(),
			Metadata:   (*e.GetMetadata()).(map[string]string),
			// Groups work in progress
		})
	}

	nodes := make([]*neocortex.DialogNode, 0)

	pp.Println(res.Output.GetNodesVisitedDetails())

	res.Output.GetNodesVisitedDetails()
	for _, n := range *res.Output.GetNodesVisitedDetails() {
		nodes = append(nodes, &neocortex.DialogNode{
			Name:       *n.DialogNode,
			Conditions: *n.Conditions,
			Title:      *n.Title,
		})

	}

	genResponses := make([]*neocortex.ResponseGeneric, 0)

	for _, r := range *(res.Output.GetGeneric()) {
		genResponses = append(genResponses, &neocortex.ResponseGeneric{
			Text:                *r.Text,
			Typing:              *r.Typing,
			Source:              *r.Source,
			Time:                time.Unix(*r.Time, 0),
			Topic:               *r.Topic,
			Description:         *r.Description,
			ResponseType:        neocortex.ResponseGenericType(*r.ResponseType),
			MessageToHumanAgent: *r.MessageToHumanAgent,
			Preference:          *r.Preference,
		})
	}

	return &neocortex.Output{
		Context:         context,
		Intents:         intents,
		Entities:        entities,
		InputText:       *res.Input.Text,
		NodesVisited:    nodes,
		OutputText:      *res.Output.GetText(),
		GenericResponse: []*neocortex.ResponseGeneric{},
		// TODO: Actions
		// TODO: Logs
	}
}

func (watson *WatsonCognitiveService) GetProtoResponse(in *neocortex.Input) (*neocortex.Output, error) {
	opts := watson.getNativeIn(in)
	pp.Println(opts)
	r, err := watson.assistant.Message(opts)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, errors.New("err0xF3")
	}

	response := watson.assistant.GetMessageResult(r)

	pp.Println(response)
	out := watson.getNeoCortexOut(response)
	return out, nil
}

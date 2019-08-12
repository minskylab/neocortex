package neocortex

import (
	"sort"
	"time"
)

type MessageOwner string

const Person MessageOwner = "person"
const ChatBot MessageOwner = "bot"

type MessageResponse struct {
	Type  ResponseType
	Value interface{}
}

type Message struct {
	At       time.Time
	Owner    MessageOwner
	Intents  []Intent
	Entities []Entity
	Response MessageResponse
}

type Chat struct {
	ID            string
	LastMessageAt time.Time
	Person        PersonInfo
	Performance   float64
	Messages      []Message
}

type byDate []Message

func (messages byDate) Len() int {
	return len(messages)
}
func (messages byDate) Swap(i, j int) {
	messages[i], messages[j] = messages[j], messages[i]
}
func (messages byDate) Less(i, j int) bool {
	return messages[i].At.Before(messages[j].At)
}

type byLastMessageAt []*Chat

func (chats byLastMessageAt) Len() int {
	return len(chats)
}
func (chats byLastMessageAt) Swap(i, j int) {
	chats[i], chats[j] = chats[j], chats[i]
}
func (chats byLastMessageAt) Less(i, j int) bool {
	return chats[i].LastMessageAt.After(chats[j].LastMessageAt)
}

func (analitycs *Analytics) processDialogs(dialogs []*Dialog) []*Chat {
	performances := map[string][]float64{}
	mapChats := map[string]*Chat{}
	lastMessages := map[string]time.Time{}

	for _, d := range dialogs {
		if len(d.Contexts) == 0 {
			continue
		}

		ctx := d.Contexts[0].Context
		personID := ctx.Person.ID
		if _, ok := mapChats[personID]; !ok {
			mapChats[personID] = new(Chat)
			mapChats[personID].Person = ctx.Person
			mapChats[personID].ID = personID
			mapChats[personID].Messages = []Message{}

		}

		if _, ok := performances[personID]; !ok {
			performances[personID] = []float64{}
		}

		if _, ok := lastMessages[personID]; !ok {
			lastMessages[personID] = time.Time{}
		}

		for _, i := range d.Ins {
			mapChats[personID].Messages = append(mapChats[personID].Messages, Message{
				At:    i.At,
				Owner: Person,
				Response: MessageResponse{
					Type:  Text,
					Value: i.Input.Data.Value,
				},
				Entities: i.Input.Entities,
				Intents:  i.Input.Intents,
			})
			if i.At.After(lastMessages[personID]) {
				lastMessages[personID] = i.At
			}
		}

		for _, o := range d.Outs {
			for _, r := range o.Output.Responses {
				mapChats[personID].Messages = append(mapChats[personID].Messages, Message{
					At:    o.At,
					Owner: ChatBot,
					Response: MessageResponse{
						Type:  r.Type,
						Value: r.Value,
					},
					Entities: o.Output.Entities,
					Intents:  o.Output.Intents,
				})
			}

		}

		if analitycs.performanceFunction != nil {
			performances[personID] = append(performances[personID], analitycs.performanceFunction(d))
		}

	}

	chats := make([]*Chat, 0)
	for personID, chat := range mapChats {
		performance := 0.0
		for _, p := range performances[personID] {
			performance += p
		}

		performance /= float64(len(performances[personID]))

		sort.Sort(byDate(chat.Messages))

		chats = append(chats, &Chat{
			ID:            chat.ID,
			Person:        chat.Person,
			Messages:      chat.Messages,
			Performance:   performance,
			LastMessageAt: lastMessages[personID],
		})
	}

	sort.Sort(byLastMessageAt(chats))

	return chats
}

type TimeAnalysisResult struct {
	Timeline    map[time.Time]map[string]float64
	TotalCounts int
}

func (analitycs *Analytics) timeAnalysis(viewID string, frame TimeFrame) (*TimeAnalysisResult, error) {
	view, err := analitycs.repo.GetViewByID(viewID)
	if err != nil {
		return nil, err
	}

	dialogs, err := analitycs.repo.DialogsByView(viewID, frame)
	if err != nil {
		return nil, err
	}

	timeline := map[time.Time]map[string]float64{}
	totalCounts := 0
	for _, dialog := range dialogs {
		for _, class := range view.Classes {
			valueName := ""
			switch class.Type {
			case EntityClass:
				if dialog.HasEntity(class.Value) {
					valueName = "E-" + class.Value
				}
			case IntentClass:
				if dialog.HasIntent(class.Value) {
					valueName = "I-" + class.Value
				}
			case DialogNodeClass:
				if dialog.HasDialogNode(class.Value) {
					valueName = "D-" + class.Value
				}
			case ContextVarClass:
				// TODO: Implement that
				continue
			default:
				continue
			}

			if _, ok := timeline[dialog.StartAt]; !ok {
				timeline[dialog.StartAt] = map[string]float64{}
			}
			if _, ok := timeline[dialog.StartAt][valueName]; !ok {
				timeline[dialog.StartAt][valueName] = 0.0
			}
			timeline[dialog.StartAt][valueName]++
			totalCounts++
		}
	}

	return &TimeAnalysisResult{
		Timeline:    timeline,
		TotalCounts: totalCounts,
	}, nil

}

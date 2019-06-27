package mongodb

import (
	"time"

	"github.com/k0kubun/pp"

	"github.com/bregydoc/neocortex"
	"go.mongodb.org/mongo-driver/bson"
)

type DialogDocument struct {
	ID       string    `bson:"id"`
	StartAt  time.Time `bson:"start_at"`
	EndAt    time.Time `bson:"end_at"`
	Ins      bson.M    `bson:"ins"`
	Outs     bson.M    `bson:"outs"`
	Contexts bson.M    `bson:"contexts"`
}

func decodeIntent(intent bson.M) neocortex.Intent {
	i, _ := intent["intent"].(string)
	c, _ := intent["confidence"].(float64)

	return neocortex.Intent{
		Intent:     i,
		Confidence: c,
	}
}

func decodeEntity(entity bson.M) neocortex.Entity {
	e, _ := entity["entity"].(string)
	location, _ := entity["location"].([]int64)
	value, _ := entity["value"].(string)
	confidence, _ := entity["confidence"].(float64)
	metadata, _ := entity["metadata"].(map[string]interface{})

	return neocortex.Entity{
		Entity:     e,
		Location:   location,
		Value:      value,
		Confidence: confidence,
		Metadata:   metadata,
	}
}

func decodeIntents(intents []bson.M) []neocortex.Intent {
	list := make([]neocortex.Intent, 0)
	for _, i := range intents {
		intent := decodeIntent(i)
		list = append(list, intent)
	}

	return list
}

func decodeEntities(entities []bson.M) []neocortex.Entity {
	list := make([]neocortex.Entity, 0)
	for _, e := range entities {
		entity := decodeEntity(e)
		list = append(list, entity)
	}

	return list
}

func decodeInputData(inputData bson.M) neocortex.InputData {
	t, _ := inputData["type"].(string)
	v, _ := inputData["value"].(string)
	d, _ := inputData["data"].(string)

	return neocortex.InputData{
		Type:  neocortex.InputType(t),
		Value: v,
		Data:  []byte(d),
	}
}

func decodeInput(in bson.M) neocortex.Input {
	// inData := InputData {
	// 	Type  InputType `json:"type"`
	// 	Value string    `json:"value"`
	// 	Data  []byte    `json:"data"`
	// }
	tt, _ := in["data"].(bson.M)
	ens, _ := in["entities"].([]bson.M)
	ins, _ := in["intents"].([]bson.M)

	entities := decodeEntities(ens)
	intents := decodeIntents(ins)
	inType := decodeInputData(tt)

	return neocortex.Input{
		Data:     inType,
		Entities: entities,
		Intents:  intents,
	}
}

func dialogToDocument(dialog *neocortex.Dialog) *DialogDocument {
	ins := bson.M{}
	for k, i := range dialog.Ins {
		ins[k.String()] = i
	}

	outs := bson.M{}
	for k, o := range dialog.Outs {
		outs[k.String()] = o
	}

	contexts := bson.M{}
	for k, c := range dialog.Contexts {
		contexts[k.String()] = c
	}

	doc := &DialogDocument{
		ID:       dialog.ID,
		StartAt:  dialog.StartAt,
		EndAt:    dialog.EndAt,
		Ins:      ins,
		Outs:     outs,
		Contexts: contexts,
	}

	return doc
}

func documentToDialog(doc *DialogDocument) *neocortex.Dialog {
	pp.Println(doc)
	ins := map[time.Time]neocortex.Input{}
	for k, i := range doc.Ins {
		t, _ := time.Parse(time.RFC3339, k)
		ii, _ := i.(bson.M)
		in := decodeInput(ii)
		ins[t] = in
	}

	outs := map[time.Time]neocortex.Output{}
	for k, o := range doc.Outs {
		if out, ok := o.(neocortex.Output); ok {
			t, _ := time.Parse(time.RFC3339, k)
			outs[t] = out
		}
	}

	contexts := map[time.Time]neocortex.Context{}
	for k, c := range doc.Contexts {
		if context, ok := c.(neocortex.Context); ok {
			t, _ := time.Parse(time.RFC3339, k)
			contexts[t] = context
		}
	}

	return &neocortex.Dialog{
		ID:       doc.ID,
		StartAt:  doc.StartAt,
		EndAt:    doc.EndAt,
		Ins:      ins,
		Outs:     outs,
		Contexts: contexts,
	}
}

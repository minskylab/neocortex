package mongodb

import (
	"time"

	"github.com/bregydoc/neocortex"
	"go.mongodb.org/mongo-driver/bson"
)

type DialogDocument struct {
	ID       string             `bson:"id"`
	Context  *neocortex.Context `bson:"context"`
	StartAt  time.Time          `bson:"start_at"`
	EndAt    time.Time          `bson:"end_at"`
	Ins      bson.M             `bson:"ins"`
	Outs     bson.M             `bson:"outs"`
	Contexts bson.M             `bson:"contexts"`
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
	ins := map[time.Time]neocortex.Input{}
	for k, i := range doc.Ins {
		t, _ := time.Parse(time.RFC3339, k)
		in, ok := i.(neocortex.Input)
		if ok {
			ins[t] = in
		}
	}

	outs := map[time.Time]neocortex.Output{}
	for k, o := range doc.Outs {
		t, _ := time.Parse(time.RFC3339, k)
		out, ok := o.(neocortex.Output)
		if ok {
			outs[t] = out
		}
	}

	contexts := map[time.Time]neocortex.Context{}
	for k, c := range doc.Contexts {
		t, _ := time.Parse(time.RFC3339, k)
		context, ok := c.(neocortex.Context)
		if ok {
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

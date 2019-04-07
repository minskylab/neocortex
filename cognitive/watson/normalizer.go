package watson

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

func getNativeEntity(e *neo.Entity) assistantv2.RuntimeEntity {
	return assistantv2.RuntimeEntity{
		Confidence: &e.Confidence,
		Entity:     &e.Entity,
		Location:   e.Location,
		Value:      &e.Value,
		Metadata:   e.Metadata,
	}
}

func getNativeIntent(i *neo.Intent) assistantv2.RuntimeIntent {
	return assistantv2.RuntimeIntent{
		Intent:     &i.Intent,
		Confidence: &i.Confidence,
	}
}

func getNeocortexIntent(i assistantv2.RuntimeIntent) neo.Intent {
	return neo.Intent{
		Intent:     *i.Intent,
		Confidence: *i.Confidence,
	}
}
func getNeocortexEntity(i assistantv2.RuntimeEntity) neo.Entity {
	metadata, ok := i.Metadata.(map[string]interface{})
	if !ok {
		metadata = map[string]interface{}{}
	}
	return neo.Entity{
		Entity:     *i.Entity,
		Metadata:   metadata,
		Confidence: *i.Confidence,
		Location:   i.Location,
		Value:      *i.Value,
	}
}

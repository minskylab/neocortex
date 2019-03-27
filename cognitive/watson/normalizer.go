package watson

import (
	"github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

func getNativeEntity(e *neocortex.Entity) assistantv2.RuntimeEntity {
	return assistantv2.RuntimeEntity{
		Confidence: &e.Confidence,
		Entity:     &e.Entity,
		Location:   e.Location,
		Value:      &e.Value,
		Metadata:   e.Metadata,
	}
}

func getNativeIntent(i *neocortex.Intent) assistantv2.RuntimeIntent {
	return assistantv2.RuntimeIntent{
		Intent:     &i.Intent,
		Confidence: &i.Confidence,
	}
}

func getNeocortexIntent(i assistantv2.RuntimeIntent) neocortex.Intent {
	return neocortex.Intent{
		Intent:     *i.Intent,
		Confidence: *i.Confidence,
	}
}
func getNeocortexEntity(i assistantv2.RuntimeEntity) neocortex.Entity {
	return neocortex.Entity{
		Entity:     *i.Entity,
		Metadata:   i.Metadata.(map[string]interface{}),
		Confidence: *i.Confidence,
		Location:   i.Location,
		Value:      *i.Value,
	}
}

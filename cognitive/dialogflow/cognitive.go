package dialogflow

import (
	"errors"
	"reflect"
)

type Cognitive struct {
	tokenIdentify string
	Url           string
}

type NewCognitiveParams struct {
	Url         string
	AccessToken string
	Version     string
	sessionId   string
}

func NewCognitive(params NewCognitiveParams) (*Cognitive, error) {
	if (params.AccessToken == "" || reflect.DeepEqual(params, NewCognitiveParams{})) {
		return nil, errors.New("No token found")
	}

	client := &Cognitive{
		Url: params.Url,
	}

	return client, nil
}

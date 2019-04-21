package main

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/bregydoc/neocortex/channels/terminal"
	"github.com/bregydoc/neocortex/cognitive/watson"
	"github.com/bregydoc/neocortex/repositories/boltdb"
)

func main() {
	watsonAgent, err := watson.NewCognitive(watson.NewCognitiveParams{
		Url:         "<WATSON_URL_API>",
		Username:    "<WATSON_USERNAME",
		Password:    "<WATSON_PASSWORD>",
		Version:     "<WATSON_VERSION>",
		AssistantID: "<ASSISTANT_ID>",
	})
	if err != nil {
		panic(err)
	}

	term := terminal.NewChannel(nil)

	repo, err := boltdb.New("neocortex.db")
	if err != nil {
		panic(err)
	}

	engine, err := neo.Default(repo, watsonAgent, fb)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	engine.ResolveAny(term, func(in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		return response(out)
	})

	if err := engine.Run(); err != nil {
		panic(err)
	}
}

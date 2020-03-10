package main

import (
	neo "github.com/minskylab/neocortex"
	"github.com/minskylab/neocortex/channels/terminal"
	"github.com/minskylab/neocortex/cognitive/watson"
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

	// repo, err := boltdb.New("neocortex.db")
	// if err != nil {
	// 	panic(err)
	// }

	engine, err := neo.Default(nil, watsonAgent, term)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	engine.ResolveAny(term, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		return response(c, out)
	})

	if err := engine.Run(); err != nil {
		panic(err)
	}
}

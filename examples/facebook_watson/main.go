package main

import (
	neo "github.com/minskylab/neocortex"
	"github.com/minskylab/neocortex/channels/facebook"
	"github.com/minskylab/neocortex/cognitive/watson"
)

func main() {
	watsonAgent, err := watson.NewCognitive(watson.NewCognitiveParams{
		Url:         "<Url>",
		ApiKey:      "<KEY>",
		Version:     "<DATE-VERSION>",
		AssistantID: "<ID>",
	})
	if err != nil {
		panic(err)
	}

	fb, err := facebook.NewChannel(facebook.ChannelOptions{
		AccessToken: "<TOKEN>",
		VerifyToken: "<YourToken>",
		PageID:      "<PAGE-ID>",
	})
	if err != nil {
		panic(err)
	}

	// repo, err := boltdb.New("neocortex.db")
	// if err != nil {
	// 	panic(err)
	// }

	engine, err := neo.Default(nil, watsonAgent, fb)

	if err != nil {
		panic(err)
	}

	match := neo.IntentIs("HELLO")
	engine.Resolve(fb, match, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		out.Fill(map[string]string{
			"Name": c.Person.Name,
		})
		return response(c, out)
	})

	engine.ResolveAny(fb, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		out.AddTextResponse("[Unhandled]")
		return response(c, out)
	})

	err = engine.Run()
	if err != nil {
		panic(err)
	}
}

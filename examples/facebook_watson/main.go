package main

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/bregydoc/neocortex/channels/facebook"
	"github.com/bregydoc/neocortex/cognitive/watson"
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

	fb, err := facebook.NewChannel(facebook.ChannelOptions{
		AccessToken: "<Your ACCESS_TOKEN>",
		VerifyToken: "<Your VERIFY_TOKEN>",
		PageID:      "<Your PAGE_ID>",
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

	match := neo.IfIntentIs("HELLO")
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

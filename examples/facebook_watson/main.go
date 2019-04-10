package main

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/bregydoc/neocortex/channels/facebook"
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

	fb, err := facebook.NewChannel(facebook.ChannelOptions{
		AccessToken: "<Your ACCESS_TOKEN>",
		VerifyToken: "<Your VERIFY_TOKEN>",
		PageID:      "<Your PAGE_ID>",
	})
	if err != nil {
		panic(err)
	}

	repo, err := boltdb.New("neocortex.db")
	if err != nil {
		panic(err)
	}

	engine, err := neo.New(repo, watsonAgent, fb)

	if err != nil {
		panic(err)
	}

	match := neo.Matcher{
		Intent: neo.Match{Is: "HELLO", Confidence: 0.8},
	}

	engine.Resolve(fb, match, func(in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		out.Fill(map[string]string{
			"Name": in.Context.Person.Name,
		})
		return response(out)
	})

	engine.ResolveAny(fb, func(in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		out.AddTextResponse("[Unhandled]")
		return response(out)
	})

	err = engine.Run()
	if err != nil {
		panic(err)
	}
}

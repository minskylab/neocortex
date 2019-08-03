package main

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/bregydoc/neocortex/channels/facebook"
	"github.com/bregydoc/neocortex/cognitive/watson"
)

func main() {
	watsonAgent, err := watson.NewCognitive(watson.NewCognitiveParams{
		Url:         "https://gateway.watsonplatform.net/assistant/api",
		ApiKey:      "JXsq7CP7xseGQgjetozUWUeAKMq0vHIQchM7yo8svTCh",
		Version:     "2019-07-01",
		AssistantID: "cf18332a-78c1-44d7-bae7-61bc969645f7",
	})
	if err != nil {
		panic(err)
	}

	fb, err := facebook.NewChannel(facebook.ChannelOptions{
		AccessToken: "EAAhuwjZBO2tUBAGbJdfUfoD6lLSjpiZBxzUmtsKH5prHRxjJChZBZARpJbnALAEeGcOhdErOybao4QfnxXrCoVhNxxtZAZBUccj1hr2QJFYBOZCs5ABL3uVcMnAZBlEVMVqwZA2H2UH8hQ4SHdkZCNGBlwqp59qY6yaKymUMvJlojwKQZDZD",
		VerifyToken: "toche",
		PageID:      "2373580382722773",
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

package main

import (
	neo "github.com/minskylab/neocortex"
	"github.com/minskylab/neocortex/channels/facebook"

	"github.com/minskylab/neocortex/cognitive/uselessbox"
)

func main() {
	box := uselessbox.NewCognitive()

	fb, err := facebook.NewChannel(facebook.ChannelOptions{
		AccessToken: "<Your ACCESS_TOKEN>",
		VerifyToken: "<Your VERIFY_TOKEN>",
		PageID:      "<Your PAGE_ID>",
	})

	// repo, err := boltdb.New("neocortex.db")
	// if err != nil {
	// 	panic(err)
	// }

	engine, err := neo.Default(nil, box, fb)

	if err != nil {
		panic(err)
	}

	engine.ResolveAny(fb, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		return response(c, out)
	})

	if err := engine.Run(); err != nil {
		panic(err)
	}
}

package main

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/bregydoc/neocortex/channels/facebook"

	"github.com/bregydoc/neocortex/cognitive/uselessbox"
)

func main() {
	box := uselessbox.NewCognitive()

	fb, err := facebook.NewChannel(facebook.ChannelOptions{
		AccessToken: "<Your ACCESS_TOKEN>",
		VerifyToken: "<Your VERIFY_TOKEN>",
		PageID:      "<Your PAGE_ID>",
	})

	engine, err := neo.New(box, fb)

	if err != nil {
		panic(err)
	}

	engine.ResolveAny(fb, func(in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		return response(out)
	})

	if err := engine.Run(); err != nil {
		panic(err)
	}
}

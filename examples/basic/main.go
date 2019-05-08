package main

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/bregydoc/neocortex/channels/terminal"
	"github.com/bregydoc/neocortex/cognitive/uselessbox"
	"github.com/bregydoc/neocortex/repositories/boltdb"
)

// Example of use useless box with terminal channel
func main() {
	box := uselessbox.NewCognitive()
	term := terminal.NewChannel(nil)

	repo, err := boltdb.New("neocortex.db")
	if err != nil {
		panic(err)
	}

	engine, err := neo.Default(repo, box, term)
	if err != nil {
		panic(err)
	}

	engine.InjectAll(term, func(in *neo.Input) *neo.Input {
		in.Context.Variables["user_name"] = "Bregy"
		if in.Context.Variables["count"] == nil {
			in.Context.Variables["count"] = 0
		}

		in.Context.Variables["count"] = in.Context.Variables["count"].(int) + 1

		return in
	})

	engine.ResolveAny(term, func(in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		out.AddTextResponse("-----Watermark-----")
		return response(out)
	})

	if err = engine.Run(); err != nil {
		panic(err)
	}
}

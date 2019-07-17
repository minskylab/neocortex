package main

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/bregydoc/neocortex/channels/terminal"
	"github.com/bregydoc/neocortex/cognitive/uselessbox"
	"github.com/bregydoc/neocortex/repositories/mongodb"
)

// Example of use useless box with terminal channel
func main() {
	box := uselessbox.NewCognitive()
	term := terminal.NewChannel(nil)

	repo, err := mongodb.New("mongodb+srv://amanda:LZlt2PQrqJW5r5RN@amanda-520ju.mongodb.net/test?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}

	engine, err := neo.Default(repo, box, term)
	if err != nil {
		panic(err)
	}

	engine.RegisterAdmin("admin", "admin")
	engine.RegisterAdmin("antonio", "admin")
	engine.RemoveAdmin("admin")

	engine.InjectAll(term, func(c *neo.Context, in *neo.Input) *neo.Input {
		c.Variables["user_name"] = "Bregy"
		if c.Variables["count"] == nil {
			c.Variables["count"] = 0
		}

		c.Variables["count"] = c.Variables["count"].(int) + 1

		return in
	})

	engine.ResolveAny(term, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		out.AddTextResponse("-----Watermark-----")

		return response(c, out)
	})

	if err = engine.Run(); err != nil {
		panic(err)
	}
}

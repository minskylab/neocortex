package main

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/bregydoc/neocortex/channels/terminal"
	"github.com/bregydoc/neocortex/cognitive/uselessbox"
	"github.com/bregydoc/neocortex/repositories/boltdb"
	"github.com/k0kubun/pp"
)

// Example of use useless box with terminal channel
func main() {
	box := uselessbox.NewCognitive()
	term := terminal.NewChannel(nil)

	repo, err := boltdb.New("neocortex.db")
	if err != nil {
		panic(err)
	}

	engine, err := neo.New(repo, box, term)
	if err != nil {
		panic(err)
	}

	d, err := repo.GetAllDialogs()
	if err != nil {
		panic(err)
	}

	pp.Println(d[0])

	engine.ResolveAny(term, func(in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		out.AddTextResponse("-----Watermark-----")
		return response(out)
	})

	if err = engine.Run(); err != nil {
		panic(err)
	}
}

package neocortex

import (
	"context"
	"fmt"
)

type HandleResolver func(c context.Context, in *Input, out *Output) (*Output, error)

type CortexMiddleware struct {
	Brain               CognitiveService
	Channel             CommunicationChannel
	RegisteredResolvers map[string]*HandleResolver
}

func (cortex *CortexMiddleware) onMessage(in *Input, out *Output) {
	for _, n := range in.NodesVisited {
		c := context.WithValue(context.Background(), "metadata", in.Context)
		go func(c *context.Context) {
			f, ok := cortex.RegisteredResolvers[n.Name]
			if !ok {
				err := cortex.Channel.sendResponse(*c, &Output{
					OutputText: []string{"unimplemented smart response"},
				})
				if err != nil {
					panic(err)
				}
			} else {
				ff := *f
				out, err := ff(*c, in, out)
				if err != nil {
					panic(err)
				}
				err = cortex.Channel.sendResponse(*c, out)
				if err != nil {
					panic(err)
				}
			}

		}(&c)
	}
}

func (cortex *CortexMiddleware) Resolver(node *DialogNode, handler HandleResolver) {
	if cortex.RegisteredResolvers == nil {
		cortex.RegisteredResolvers = map[string]*HandleResolver{}
	}
	cortex.RegisteredResolvers[node.Name] = &handler
}

func (cortex *CortexMiddleware) When(node *DialogNode) {
	fmt.Println(node.Title)
}

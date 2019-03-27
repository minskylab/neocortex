package neocortex

import (
	"context"
	"fmt"
	"github.com/k0kubun/pp"
)

type HandleResolver func(in *Input, out *Output, response func(output *Output) error) error

type CortexMiddleware struct {
	Cognitive           CognitiveService
	Channel             CommunicationChannel
	registeredResolvers map[string]*HandleResolver
}

func NewCortex(cognitive CognitiveService, channel CommunicationChannel) (*CortexMiddleware, error) {
	middle := &CortexMiddleware{}
	middle.Channel = channel
	middle.Cognitive = cognitive
	middle.registeredResolvers = map[string]*HandleResolver{}
	err := channel.RegisterMessageEndpoint(func(message *Input, response func(output *Output) error) error {
		return middle.onMessage(message, response)
	})
	if err != nil {
		return nil, err
	}
	return middle, nil
}

func (cortex *CortexMiddleware) onMessage(in *Input, response func(output *Output) error) error {
	c := context.WithValue(context.Background(), "metadata", in.Context)
	go func(c *context.Context) {
		pp.Println(in.Intents[0].Intent)
		f, ok := cortex.registeredResolvers[in.Intents[0].Intent]
		if !ok {
			err := out(&Output{
				Response: []*ResponseGeneric{{
					Text: "unimplemented smart response",
				}},
			})
			if err != nil {
				panic(err)
			}
		} else {
			ff := *f
			ff(in, out)
			if err != nil {
				panic(err)
			}
			err = out(finalOut)
			if err != nil {
				panic(err)
			}
		}

	}(&c)

	return nil
}

func (cortex *CortexMiddleware) Resolver(node *DialogNode, handler HandleResolver) {
	if cortex.registeredResolvers == nil {
		cortex.registeredResolvers = map[string]*HandleResolver{}
	}
	cortex.registeredResolvers[node.Title] = &handler
}

func (cortex *CortexMiddleware) When(node *DialogNode) {
	fmt.Println(node.Title)
}

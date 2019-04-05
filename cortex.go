package neocortex

import (
	"fmt"
)

type OutputResponse func(output *Output) error
type HandleResolver func(in *Input, out *Output, response OutputResponse) error
type MiddleHandler func(message *Input, response OutputResponse) error
type ContextFabric func(userID string) *Context

type CortexMiddleware struct {
	Cognitive           CognitiveService
	Channel             []CommunicationChannel
	registeredResolvers map[Matcher]*HandleResolver
	resolver            *HandleResolver
}

func NewCortex(cognitive CognitiveService, channel ...CommunicationChannel) (*CortexMiddleware, error) {
	middle := &CortexMiddleware{}
	middle.Channel = channel
	middle.Cognitive = cognitive
	middle.registeredResolvers = map[Matcher]*HandleResolver{}
	for _, ch := range channel {
		err := ch.RegisterMessageEndpoint(func(message *Input, response OutputResponse) error {
			return middle.onMessage(message, response)
		})
		if err != nil {
			return nil, err
		}

		go func() {
			err := ch.ToHear()
			if err != nil {
				panic(err)
			}
		}()
	}

	return middle, nil
}

func (cortex *CortexMiddleware) onMessage(in *Input, response OutputResponse) error {
	out, err := cortex.Cognitive.GetProtoResponse(in)
	if err != nil {
		if err == ErrSessionNotExist {
			panic(err)
		}
	}

	exist := false
	for m, resolver := range cortex.registeredResolvers {
		if match(out, &m) {
			if err = (*resolver)(in, out, response); err != nil {
				return err
			}
			exist = true
		}

	}

	if cortex.resolver != nil && !exist {
		if err = (*cortex.resolver)(in, out, response); err != nil {
			return err
		}
	}

	return nil
}

func (cortex *CortexMiddleware) ResolverAll(handler HandleResolver) {
	cortex.resolver = &handler
}

func (cortex *CortexMiddleware) Resolver(matcher Matcher, handler HandleResolver) {
	if cortex.registeredResolvers == nil {
		cortex.registeredResolvers = map[Matcher]*HandleResolver{}
	}
	cortex.registeredResolvers[matcher] = &handler
}

func (cortex *CortexMiddleware) When(node *DialogNode) {
	fmt.Println(node.Title)
}

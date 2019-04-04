package neocortex

import "fmt"

type OutputResponse func(output *Output) error
type HandleResolver func(in *Input, out *Output, response OutputResponse) error
type MiddleHandler func(message *Input, response OutputResponse) error
type ContextFabric func(userID string) *Context

type CortexMiddleware struct {
	Context             *Context
	Cognitive           CognitiveService
	Channel             []CommunicationChannel
	registeredResolvers map[string]*HandleResolver
	resolver            *HandleResolver
}

func NewCortex(c *Context, cognitive CognitiveService, channel ...CommunicationChannel) (*CortexMiddleware, error) {
	middle := &CortexMiddleware{}
	middle.Channel = channel
	middle.Cognitive = cognitive
	middle.registeredResolvers = map[string]*HandleResolver{}
	middle.Context = c
	for _, ch := range channel {
		err := ch.ToHear()
		if err != nil {
			return nil, err
		}
		err = ch.RegisterMessageEndpoint(func(message *Input, response OutputResponse) error {
			fmt.Println(message.InputType.Value)
			return middle.onMessage(ch, message, response)
		})
		if err != nil {
			return nil, err
		}
	}

	return middle, nil
}

func (cortex *CortexMiddleware) onMessage(chanel CommunicationChannel, in *Input, response OutputResponse) error {
	out, err := cortex.Cognitive.GetProtoResponse(in)
	if err != nil {
		if err == ErrSessionNotExist {
			panic(err)
		}
	}

	if len(out.Intents) == 0 {
		if cortex.resolver != nil {
			resolver := *cortex.resolver
			err = resolver(in, out, response)
			if err != nil {
				panic(err)
			}
			return nil
		}
		err := response(out)
		if err != nil {
			panic(err)
		}

		return nil
	}

	f, ok := cortex.registeredResolvers[out.Intents[0].Intent]
	if !ok {
		if cortex.resolver != nil {
			resolver := *cortex.resolver
			err = resolver(in, out, response)
			if err != nil {
				panic(err)
			}
		} else {
			// out := CreateNewGenericOutputText(cortex.Context, "unimplemented smart response", nil, nil)
			err := response(out)
			if err != nil {
				panic(err)
			}
		}

	} else {
		ff := *f
		err = ff(in, out, response)
		if err != nil {
			panic(err)
		}

	}

	return nil
}

func (cortex *CortexMiddleware) ResolverAll(handler HandleResolver) {
	cortex.resolver = &handler
}

func (cortex *CortexMiddleware) Resolver(entity string, handler HandleResolver) {
	if cortex.registeredResolvers == nil {
		cortex.registeredResolvers = map[string]*HandleResolver{}
	}
	cortex.registeredResolvers[entity] = &handler
}

func (cortex *CortexMiddleware) When(node *DialogNode) {
	fmt.Println(node.Title)
}

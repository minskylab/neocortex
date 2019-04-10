# Neocortex  🧠

Neocortex is a tool to connect your cognitive service with your services and communication channels.

The main goal of neocortex is offer a reliable and modern api to connect any kind of cognitive service* with any communication channel**. 

*Currently neocortex offers only two cognitive services: Watson and a simple Uselessbox as dummy service, you can collaborate to implement another cognitive services like DialogFlow or Amazon Lex, later I'm going to documment how to implement this services but you can read the source code to understand how to.

**Like cognitive services, I could only implement only two channels: Facebook Messenger and a simple Terminal chat (very simple to emulate a chat in your terminal), if you want you can collaborate implementing another channels like Slack, Whatsapp or Gmail, for example.

*Neocortex is work in progress, it pretends to be a big collaboratibe project*

###  TODO

- [x] Describe a Cognitive Service interface (`type CognitiveService interface`)

- [x] Describe a Communication channel interface (`type CommunicationChannel interface`)

- [x] Implement the Watson cognitive service 

- [x] Implement the Facebook channel

- [ ] Write unit tests

- [ ] Make a iteration of the Communication channel's architecture 

- [ ] Think more in the Cognitive's orientation (paradigm, architecture, etc)

- [ ] Improve the neocortex engine

- [ ] Write the Gmail channel implementation

- [ ] Write the Dialogflow service implementation

- [ ] Improve facebook messenger API

- [ ] Document more

- [ ] Document more more!

  

## Install

Install with:

```go get -u github.com/bregydoc/neocortex```

Currently neocortex have 2 implementations of Cognitive Services (Useless-box and Watson Assistant based on the [watson official api v2](https://github.com/watson-developer-cloud/go-sdk)) and 2 implementation of Communication Channels (Terminal based UI and Facebook Messenger forked from [Facebook-messenger API](https://github.com/mileusna/facebook-messenger) by [mileusna](https://github.com/mileusna)).

### Basic Example

```go
package main

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/bregydoc/neocortex/channels/terminal"
	"github.com/bregydoc/neocortex/cognitive/uselessbox"
)

// Example of use useless box with terminal channel
func main() {
	box := uselessbox.NewCognitive()
	term := terminal.NewChannel(nil)
	
	engine, err := neo.New(box, term)
	if err != nil {
		panic(err)
	}

	engine.ResolveAny(term, func(in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		out.AddTextResponse("-----Watermark-----")
		return response(out)
	})

	if err = engine.Run(); err != nil {
		panic(err)
	}
}

```

## Paradigm

Neocortex is like a middleware with a mux, with it you can catch your message input, pass to your cognitive service, inflate or modify them and response.

### Concepts

1. **Cognitive Service**

   Represents a any service that decode and find intents and entities in a human message. In neocortex this is described by a simple interface.

   ```go
   type CognitiveService interface {
      CreateNewContext(c *context.Context, info neocortex.PersonInfo) *neocortex.Context
      GetProtoResponse(in *neocortex.Input) (*neocortex.Output, error)
   }
   ```

   you can see the implementation of a [Useless box](https://github.com/bregydoc/neocortex/tree/master/cognitive/uselessbox) or [Watson Assistant](https://github.com/bregydoc/neocortex/tree/master/cognitive/watson).

2. **Communication Channel**

   A Communication Channel is any human interface where a person can to send a message and receive a response. Currently I think we need to work more in the paradigm behind Communication channels. In neocortex a communication channel is described by the following interface:

   ```go
   type CommunicationChannel interface {
      RegisterMessageEndpoint(handler neocortex.MiddleHandler) error
      ToHear() error
      GetContextFabric() neocortex.ContextFabric
      SetContextFabric(fabric neocortex.ContextFabric)
      OnNewContextCreated(callback func(c *neocortex.Context))
   }
   ```

   Please, read how to are implemented the [Terminal channel](https://github.com/bregydoc/neocortex/tree/master/channels/terminal) or [Facebook Messenger Channel](https://github.com/bregydoc/neocortex/tree/master/channels/facebook).

3. **Context**

   Neocortex's Context represents a "session" or "dialog" with a human, it contains essencial information about the person with we're conversation.

   ```go
   // Context represent the context of one conversation
   type Context struct {
      Context   *context.Context // go native context implementation
      SessionID string           
      Person    neocortex.PersonInfo             
      Variables map[string]interface{} //conversation context variables
   }
   ```

4. **Input**

   An input is a message input, that's all. An input have a specified type and in neocortex is a struct:

   ```go
   type Input struct {
      Context  *neocortex.Context
      Data     neocortex.InputData
      Entities []neocortex.Entity
      Intents  []neocortex.Intent
   }
   ```

   Intents and Entities define the message.

5. **Output**

   An output represents a response in a conversation, with this you can define the response to your communication channel (e.g. facebook messenger) and if your channel allow you can responde different types of response (e.g. Image, Audio,  Attachment, etc).

   ```go
   type Output struct {
      Context      *neocortex.Context
      Entities     []neocortex.Entity
      Intents      []neocortex.Intent
      VisitedNodes []*neocortex.DialogNode
      Logs         []*neocortex.LogMessage
      Responses    []neocortex.Response // A list of responses
   }
   
   type Response struct {
   	IsTyping bool
   	Type     neocortex.ResponseType
   	Value    interface{}
   }
   ```

6. **Engine**

   This is the core of neocortex it can to connect and manage your Cognitive service and your Communication channels. the engine have a different methods for intercept a message and modify it.

   ```go
   // Create a New neocortex Engine
   func New(cognitive CognitiveService, channels ...CommunicationChannel) (*Engine, error)
   
   // Register a new resolver, you need a matcher
   func (engine *Engine) Resolve(channel CommunicationChannel, matcher Matcher, handler HandleResolver)
   
   // Maatcher
   type Matcher struct {
   	Entity Match
   	Intent Match
   	AND    *Matcher
   	OR     *Matcher
   }
   // Match
   type Match struct {
   	Is         string
   	Confidence float64
   }
   ```

   

7. **Resolver**

   This is core of the neocortex paradigm, with this you can intercept message and modify or only bypass it. You need to pass a Matcher who is used to match with your message inputs, you can see below how looks like a Matcher struct. Above you can see two examples of a matcher:

   ```go
   // match if the input have a Regard or Goodbye intents
   match := neo.Matcher{Intent: neo.Match{Is: "REGARD"}, OR: &neo.Matcher{Intent: neo.Match{Is: "GOODBYE"}}}
   
   // match if the input is an insult intent and have a bad_word entity
   match := neo.Matcher{
     Intent: neo.Match{Is:"INSULT", Confidence: 0.8}, 
     AND: &neo.Matcher{
       Entity: neo.Match{
         Is: "bad_word",
       },
     },
   }
   ```

   Register a new resolver is simple, see the follow example above

   ```go
   match := neo.Matcher{Intent: neo.Match{Is:"HELLO", Confidence: 0.8}}
   engine.Resolve(fb, match, func(in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
      out.AddTextResponse("Powered by neocortex")
      return response(out)
   })
   ```

   You can make another types of Resolves.

   ```go
   func (engine *Engine) ResolveAny(channel CommunicationChannel, handler HandleResolver) 
   func (engine *Engine) Resolve(channel CommunicationChannel, matcher Matcher, handler HandleResolver)
   func (engine *Engine) ResolveMany(channels []CommunicationChannel, matcher Matcher, handler HandleResolver)
   func (engine *Engine) ResolveManyAny(channels []CommunicationChannel, handler HandleResolver)
   ```



🚧 Work in progress documentation, if you can help, only send me an email.



![love open source](https://github.com/bregydoc/torioux-hands/raw/master/I_love_opensource.png)




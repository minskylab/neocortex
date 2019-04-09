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

   // todo

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

   

5. **Output**

   // todo

6. Engine

   // todo

7. Resolver

   // todo






package terminal

import (
	"bufio"
	"fmt"
	"github.com/bregydoc/neocortex"
	"os"
	"strconv"
)

// This channel have particularity only one user, then we have only one user ID for all
const uniqueUserID = 0

type Channel struct {
	reader     *bufio.Reader
	options    *ChannelOptions
	messageIn  neocortex.MiddleHandler      // req
	newContext neocortex.ContextFabric      // req
	contexts   map[int64]*neocortex.Context // req
}

type ChannelOptions struct {
	PersonIcon string
	PersonName string
	BotIcon    string
	BotName    string
	SaysSymbol string
}

func NewChannel(factory neocortex.ContextFabric, opts *ChannelOptions) *Channel {
	if opts == nil { // default
		opts = &ChannelOptions{
			PersonIcon: "😀",
			PersonName: "User",
			BotIcon:    "🤖",
			BotName:    "Bot",
			SaysSymbol: " >",
		}
	}
	t := &Channel{
		reader:     bufio.NewReader(os.Stdin),
		options:    opts,
		contexts:   map[int64]*neocortex.Context{},
		newContext: factory,
	}

	return t
}
func (term *Channel) getInput() string {
	input, _ := term.reader.ReadString('\n')
	input = input[:len(input)-1]
	return input
}

func (term *Channel) renderUserInterface(done bool) error {
	c, contextExist := term.contexts[uniqueUserID]
	if !contextExist {
		c = term.newContext(strconv.Itoa(uniqueUserID))
		term.contexts[uniqueUserID] = c
	}
	if !done {
		fmt.Printf("%s %s[sess:%s]%s ", term.options.PersonIcon, term.options.PersonName, c.SessionID, term.options.SaysSymbol)
		inStr := term.getInput()
		input := term.NewInputText(c, inStr, nil, nil)
		err := term.messageIn(input, func(out *neocortex.Output) error {
			for _, r := range out.Responses {
				if r.Type == neocortex.Text {
					fmt.Printf("%s %s[sess:%s]%s %s\n",
						term.options.BotIcon,
						term.options.BotName,
						out.Context.SessionID,
						term.options.SaysSymbol,
						r.Value.(string),
					)
				}
			}
			return nil
		})

		if err != nil {
			err := term.renderUserInterface(true)
			if err != nil {
				return err
			}
			return nil
		}

		err = term.renderUserInterface(false)
		if err != nil {
			return err
		}
		return nil
	} else {
		return nil
	}

}

func (term *Channel) RegisterMessageEndpoint(handler neocortex.MiddleHandler) error {
	term.messageIn = handler
	return nil
}

func (term *Channel) ToHear() error {
	return term.renderUserInterface(false)
}

func (term *Channel) GetContextFabric() neocortex.ContextFabric {
	return term.newContext
}

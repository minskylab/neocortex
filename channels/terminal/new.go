package terminal

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	neo "github.com/bregydoc/neocortex"
)

// This channel have particularity only one user, then we have only one user ID for all
const uniqueUserID = 0

type ChannelOptions struct {
	PersonIcon string
	PersonName string
	BotIcon    string
	BotName    string
	SaysSymbol string
}

func NewChannel(opts *ChannelOptions, fabric ...neo.ContextFabric) *Channel {
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
		reader:   bufio.NewReader(os.Stdin),
		options:  opts,
		contexts: map[int64]*neo.Context{},
	}

	if len(fabric) > 0 {
		f := fabric[0]
		t.newContext = f
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
		c = term.newContext(context.Background(), neo.PersonInfo{
			ID:   strconv.Itoa(uniqueUserID),
			Name: "Jhon Doe",
		})

		for _, call := range term.newContextCallbacks {
			(*call)(c)
		}
		term.contexts[uniqueUserID] = c
	}
	if !done {
		fmt.Printf("%s %s[sess:%s]%s ", term.options.PersonIcon, term.options.PersonName, c.SessionID, term.options.SaysSymbol)
		inStr := term.getInput()
		input := term.NewInputText(inStr, nil, nil)
		err := term.messageIn(c, input, func(c *neo.Context, out *neo.Output) error {
			for _, r := range out.Responses {
				if r.Type == neo.Text {
					fmt.Printf("%s %s[sess:%s]%s %s\n",
						term.options.BotIcon,
						term.options.BotName,
						c.SessionID,
						term.options.SaysSymbol,
						r.Value.(string),
					)
				}
			}
			return nil
		})

		if err != nil {
			return err
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

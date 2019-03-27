package channels

import (
	"bufio"
	"fmt"
	"github.com/bregydoc/neocortex"

	"os"
)

type TerminalChannel struct {
	reader    *bufio.Reader
	options   *TerminalChannelOptions
	messageIn neocortex.MiddleHandler
}

type TerminalChannelOptions struct {
	PersonIcon string
	PersonName string
	BotIcon    string
	BotName    string
	SaysSymbol string
}

func NewTerminalChannel(opts *TerminalChannelOptions) *TerminalChannel {
	if opts == nil { // default
		opts = &TerminalChannelOptions{
			PersonIcon: "😀",
			PersonName: "User",
			BotIcon:    "🤖",
			BotName:    "Bot",
			SaysSymbol: "|>",
		}
	}
	t := &TerminalChannel{
		reader:  bufio.NewReader(os.Stdin),
		options: opts,
	}

	go func() {
		err := t.renderUserInterface(false)
		if err != nil {
			panic(err)
		}
	}()

	return t
}

func (term *TerminalChannel) getInput() string {
	input, _ := term.reader.ReadString('\n')
	input = input[:len(input)-1]
	return input
}

func (term *TerminalChannel) renderUserInterface(done bool) error {
	if !done {
		fmt.Printf("%s %s%s ", term.options.PersonIcon, term.options.PersonName, term.options.SaysSymbol)
		input := term.getInput()
		err := term.messageIn(&neocortex.Input{Text: input}, func(output *neocortex.Output) error {
			for _, r := range output.Response {
				fmt.Printf("%s %s%s %s\n", term.options.BotIcon, term.options.BotName, term.options.SaysSymbol, r.Text)
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

	} else {
		return nil
	}
	err := term.renderUserInterface(false)
	if err != nil {
		return err
	}
	return nil
}

func (term *TerminalChannel) RegisterMessageEndpoint(handler neocortex.MiddleHandler) error {
	term.messageIn = handler
	return nil
}

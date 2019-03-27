package channels

import (
	"github.com/bregydoc/neocortex"
	"github.com/mileusna/facebook-messenger"

	"net/http"
)

type FacebookChannel struct {
	m         *messenger.Messenger
	messageIn neocortex.MiddleHandler
}

func CreateNewFacebookChannel(accessToken, verifyToken, pageID string) (*FacebookChannel, error) {
	channel := &FacebookChannel{}
	hook := func(msn *messenger.Messenger, userID int64, m messenger.FacebookMessage) {
		msn.SendTextMessage(userID, "I'm a man in the middle ["+m.Text+"]")
		in := &neocortex.Input{
			Text: m.Text,
		}
		err := channel.messageIn(in, func(output *neocortex.Output) error {
			for _, r := range output.Response {
				_, err := msn.SendTextMessage(userID, r.Text)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
	}

	channel.m = &messenger.Messenger{
		AccessToken:     accessToken,
		VerifyToken:     verifyToken,
		PageID:          pageID,
		MessageReceived: hook,
	}

	return channel, nil
}

func (fb *FacebookChannel) RegisterMessageEndpoint(handler neocortex.MiddleHandler) error {
	fb.messageIn = handler
	return nil
}

func (fb *FacebookChannel) LaunchAndWait() error {
	http.Handle("/fb-channel", fb.m)
	return http.ListenAndServe(":8008", nil)
}

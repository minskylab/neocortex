package facebook

import "github.com/minskylab/neocortex/channels/facebook/messenger"

func sendImageResponse(userID int64, msn *messenger.Messenger, url string) error {
	gm := msn.NewImageMessage(userID, url)
	_, err := msn.SendMessage(gm)
	return err
}

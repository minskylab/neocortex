package facebook

import "github.com/bregydoc/neocortex/channels/facebook/messenger"

func sendTextResponse(userID int64, msn *messenger.Messenger, text string) error {
	_, err := msn.SendTextMessage(userID, text)
	if err != nil {
		return err
	}
	return nil
}

package facebook

import "github.com/bregydoc/neocortex/channels/facebook/messenger"

func sendImageResponse(userID int64, msn *messenger.Messenger, url string) error {
	gm := msn.NewGenericMessage(userID)
	gm.AddNewElement("", "", "", url, nil)
	_, err := msn.SendMessage(gm)
	return err
}

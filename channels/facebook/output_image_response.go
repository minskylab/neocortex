package facebook

import "github.com/mileusna/facebook-messenger"

func sendImageResponse(userID int64, msn *messenger.Messenger, url string) error {
	gm := msn.NewGenericMessage(userID)
	gm.AddNewElement("", "", "", url, nil)
	_, err := msn.SendMessage(gm)
	return err
}

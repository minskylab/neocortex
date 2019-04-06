package facebook

import (
	"errors"
	"github.com/mileusna/facebook-messenger"
	"strconv"
	"time"
)

func sendPauseResponse(userID int64, msn *messenger.Messenger, pause interface{}) error {
	switch pause.(type) {
	case int: // in milliseconds
		delay := time.Duration(pause.(int64))
		time.Sleep(time.Millisecond * delay)
	case time.Duration:
		delay := pause.(time.Duration)
		time.Sleep(delay)
	case string: // in milliseconds
		delay, err := strconv.Atoi(pause.(string))
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * time.Duration(delay))
	default:
		return errors.New("invalid value, it cannot be parsed as duration or same")
	}
	return nil
}

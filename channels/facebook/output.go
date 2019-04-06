package facebook

import (
	"errors"
	"fmt"
	"github.com/bregydoc/neocortex"
	"github.com/mileusna/facebook-messenger"
)

func decodeOutput(userID int64, msn *messenger.Messenger, out *neocortex.Output) error {
	for _, r := range out.Responses {
		switch r.Type {
		case neocortex.Text:
			return sendTextResponse(userID, msn, r.Value.(string))
		case neocortex.Options:
			options, isOne := r.Value.(neocortex.OptionsResponse)
			optionsArray, isArray := r.Value.([]neocortex.OptionsResponse)

			if isOne && !isArray {
				return sendOneOptionResponse(userID, msn, options)
			} else if !isOne && isArray {
				return sendManyOptionsResponse(userID, msn, optionsArray)
			}

			return errors.New("invalid value, it cannot be parsed as OptionResponse struct")
		case neocortex.Pause:
			// Unsupported by facebook messenger
			// emulated with delay (it's so stupid)
			return sendPauseResponse(userID, msn, r.Value)
		case neocortex.Image:
			url, ok := r.Value.(string)
			if !ok {
				return errors.New("invalid value, it must be a string")
			}
			return sendImageResponse(userID, msn, url)
		case neocortex.Suggestion:
			// Unsupported by facebook messenger
		case neocortex.Unknown:
			// Unsupported by facebook messenger
		default:
			// by default neocortex sends a raw stringify of the value
			_, err := msn.SendTextMessage(userID, fmt.Sprintf("%v", r.Value))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

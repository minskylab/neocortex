package facebook

import (
	"errors"
	"fmt"

	"github.com/bregydoc/neocortex"
	"github.com/bregydoc/neocortex/channels/facebook/messenger"
)

func decodeOutput(userID int64, msn *messenger.Messenger, out *neocortex.Output) error {
	for _, r := range out.Responses {
		switch r.Type {
		case neocortex.Text:
			if err := sendTextResponse(userID, msn, r.Value.(string)); err != nil {
				return err
			}
		case neocortex.Options:
			options, isOne := r.Value.(neocortex.OptionsResponse)
			optionsArray, isArray := r.Value.([]neocortex.OptionsResponse)

			if isOne && !isArray {
				if err := sendOneOptionResponse(userID, msn, options); err != nil {
					return err
				}
			} else if !isOne && isArray {
				if err := sendManyOptionsResponse(userID, msn, optionsArray); err != nil {
					return err
				}
			} else {
				return errors.New("invalid value, it cannot be parsed as OptionResponse struct")
			}
			return nil

		case neocortex.Pause:
			// * Unsupported by facebook messenger
			// * emulated with delay (it's so stupid)
			if err := sendPauseResponse(userID, msn, r.Value); err != nil {
				return err
			}
		case neocortex.Image:
			if url, ok := r.Value.(string); ok {
				return sendImageResponse(userID, msn, url)
			}
			return errors.New("invalid value, it must be a string")
		case neocortex.Suggestion:
			// * Unsupported by facebook messenger
		case neocortex.Unknown:
			// * Unsupported by facebook messenger
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

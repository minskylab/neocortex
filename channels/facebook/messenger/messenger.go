package messenger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

const apiURL = "https://graph.facebook.com/v2.6/"
const urlTemplate = "https://graph.facebook.com/%d?fields=first_name,last_name,profile_pic,locale,timezone&access_token=%s"

// TestURL to mock FB server, used for testing
var TestURL = ""

type UserInfo struct {
	ID         int64
	Name       string
	Timezone   float64
	Locale     string
	ProfilePic string
}

// Messenger struct
type Messenger struct {
	AccessToken string
	VerifyToken string
	PageID      string

	apiURL  string
	pageURL string

	// MessageReceived event fires when message from Facebook received
	MessageReceived func(msng *Messenger, user UserInfo, m FacebookMessage)

	// DeliveryReceived event fires when delivery report from Facebook received
	// Omit (nil) if you don't want to manage this events
	DeliveryReceived func(msng *Messenger, user UserInfo, d FacebookDelivery)

	// PostbackReceived event fires when postback received from Facebook server
	// Omit (nil) if you don't use postbacks and you don't want to manage this events
	PostbackReceived func(msng *Messenger, user UserInfo, p FacebookPostback)
}

// New creates new messenger instance
func New(accessToken, pageID string) Messenger {
	return Messenger{
		AccessToken: accessToken,
		PageID:      pageID,
	}
}

// SendMessage sends chat message
func (msng *Messenger) SendMessage(m Message) (FacebookResponse, error) {
	if msng.apiURL == "" {
		if TestURL != "" {
			msng.apiURL = TestURL + "me/messages?access_token=" + msng.AccessToken // testing, mock FB URL
		} else {
			msng.apiURL = apiURL + "me/messages?access_token=" + msng.AccessToken
		}
	}

	s, _ := json.Marshal(m)
	req, err := http.NewRequest("POST", msng.apiURL, bytes.NewBuffer(s))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return FacebookResponse{}, err
	}

	return decodeResponse(resp)
}

// SendTextMessage sends text messate to receiverID
// it is shorthand instead of crating new text message and then sending it
func (msng Messenger) SendTextMessage(receiverID int64, text string) (FacebookResponse, error) {
	m := msng.NewTextMessage(receiverID, text)
	return msng.SendMessage(&m)
}

// ServeHTTP is HTTP handler for Messenger so it could be directly used as http.Handler
func (msng *Messenger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msng.VerifyWebhook(w, r)    // verify webhook if needed
	fbRq, _ := DecodeRequest(r) // get FacebookRequest object

	for _, entry := range fbRq.Entry {
		for _, msg := range entry.Messaging {
			userID := msg.Sender.ID

			url := fmt.Sprintf(urlTemplate, userID, msng.AccessToken)
			r, _ := http.Get(url)
			data := make(map[string]interface{})
			log.Println(r.StatusCode)
			d, _ := ioutil.ReadAll(r.Body)
			pp.Println(d)
			_ = json.NewDecoder(r.Body).Decode(&data)

			name := data["first_name"].(string) + " " + data["last_name"].(string)
			tz := data["timezone"].(float64)
			locale := data["locale"].(string)
			pic := data["profile_pic"].(string)
			user := UserInfo{
				ID:         userID,
				Name:       name,
				Timezone:   tz,
				Locale:     locale,
				ProfilePic: pic,
			}

			switch {
			case msg.Message != nil && msng.MessageReceived != nil:
				go msng.MessageReceived(msng, user, *msg.Message)

			case msg.Delivery != nil && msng.DeliveryReceived != nil:
				go msng.DeliveryReceived(msng, user, *msg.Delivery)

			case msg.Postback != nil && msng.PostbackReceived != nil:
				go msng.PostbackReceived(msng, user, *msg.Postback)
			}
		}
	}
}

// VerifyWebhook verifies your webhook by checking VerifyToken and sending challange back to Facebook
func (msng Messenger) VerifyWebhook(w http.ResponseWriter, r *http.Request) {
	// Facebook sends this query for verifying webhooks
	// hub.mode=subscribe&hub.challenge=1085525140&hub.verify_token=moj_token
	if r.FormValue("hub.mode") == "subscribe" {
		if r.FormValue("hub.verify_token") == msng.VerifyToken {
			w.Write([]byte(r.FormValue("hub.challenge")))
			return
		}
	}
}

// DecodeRequest decodes http request from FB messagner to FacebookRequest struct
// DecodeRequest will close the Body reader
// Usually you don't have to use DecodeRequest if you setup events for specific types
func DecodeRequest(r *http.Request) (FacebookRequest, error) {
	defer r.Body.Close()
	var fbRq FacebookRequest
	err := json.NewDecoder(r.Body).Decode(&fbRq)
	return fbRq, err
}

// decodeResponse decodes Facebook response after sending message, usually contains MessageID or Error
func decodeResponse(r *http.Response) (FacebookResponse, error) {
	defer r.Body.Close()
	var fbResp rawFBResponse
	err := json.NewDecoder(r.Body).Decode(&fbResp)
	if err != nil {
		return FacebookResponse{}, err
	}

	if fbResp.Error != nil {
		return FacebookResponse{}, fbResp.Error.Error()
	}

	return FacebookResponse{
		MessageID:   fbResp.MessageID,
		RecipientID: fbResp.RecipientID,
	}, nil
}

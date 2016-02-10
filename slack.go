package slack

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/stewart/slack/events"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"errors"

	"fmt"
)



// A Client maintains a WebSocket connection to the Slack RTM API, spitting out
// events on the Incoming channel as they come in.
type Client struct {
	Token     string
	Connected bool
	Incoming  chan interface{}
	Errors    chan error
	conn      *websocket.Conn

	// this gets incremented anytime we send a message
	messageID int
}

// Creates a new Client instance with the provided authentication token.
func NewClient(token string) *Client {
	return &Client{
		Token:     token,
		Incoming:  make(chan interface{}),
		Errors:    make(chan error, 1),
		messageID: 1,
	}
}

// Connects to Slack's RTM WebSocket API by requesting a connection URL via the
// `rtm.start` API method.
//
// Once complete, the Client is considered "connect" to Slack.
func (client *Client) Connect() error {
	url, err := requestRTM(client.Token)
	if err != nil {
		return err
	}

	dialer := websocket.Dialer{}
	headers := http.Header{}

	conn, _, err := dialer.Dial(url, headers)
	if err != nil {
		return err
	}

	client.conn = conn
	client.Connected = true
	return nil
}

// Wrapper around a goroutine that listens for incoming messages, parsing them
// into their correct event type, then tosses them (as an `interface{}`) into
// the client.Incoming channel.
func (client *Client) Loop() {
	conn := client.conn

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				client.Errors <- err
				return
			}

			message, err := events.Parse(msg)
			if err != nil {
				client.Errors <- err
				return
			}

			client.Incoming <- message
		}
	}()
}

// Sends a message to the provided channel, with the provided text.
func (client *Client) SendMessage(channel, text string) error {
	msg := struct {
		ID      int    `json:"id"`
		Type    string `json:"type"`
		Channel string `json:"channel"`
		Text    string `json:"text"`
	}{client.messageID, "message", channel, text}

	if err := client.conn.WriteJSON(msg); err != nil {
		return err
	}

	client.messageID++

	return nil
}


func (client *Client) SendDirectMessage(user_id string, text string) error {
	dm_id, _ := client.OpenDM(user_id)
	msg := struct {
		ID      int    `json:"id"`
		Type    string `json:"type"`
		Channel string `json:"channel"`
		Text    string `json:"text"`
	}{client.messageID, "message", dm_id, text}

	if err := client.conn.WriteJSON(msg); err != nil {
		return err
	}

	client.messageID++

	return nil
}

func (client *Client) OpenDM (user_id string) (string, error) {
	fmt.Print(user_id)
	route, err := url.Parse(endpoint + "im.open")
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("token", client.Token)
	params.Add("user", user_id)
	route.RawQuery = params.Encode()

	response, err := http.Get(route.String())
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	resp := openDMResponse{}
	json.Unmarshal(body, &resp)

	if !resp.OK {
		return "", errors.New("slack error: " + resp.Error)
	}

	return resp.Channel.ID, nil

}


type openDMResponse struct {
	OK bool `json:"ok"`
	NoOp bool `json:"no_op"`
	AlreadyOpen bool `json:"already_open"`
	Channel struct {
		   ID string `json:"id"`
	   } `json:"channel"`
	Error string `json:"error"`
}
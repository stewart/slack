package slack

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/stewart/slack/events"
)

type Error struct {
	Code    int `json:"code"`
	Message int `json:"msg"`
}

type Client struct {
	conn      *websocket.Conn
	Token     string
	Connected bool
	Incoming  chan interface{}
	Errors    chan error

	// this gets incremented anytime we send a message
	messageID int
}

// An out-going message, to be sent to Slack
type Message struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func NewClient(token string) *Client {
	return &Client{
		Token:     token,
		Incoming:  make(chan interface{}),
		Errors:    make(chan error, 1),
		messageID: 1,
	}
}

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

func (client *Client) SendMessage(channel, text string) error {
	msg := Message{
		ID:      client.messageID,
		Type:    "message",
		Channel: channel,
		Text:    text,
	}

	if err := client.conn.WriteJSON(msg); err != nil {
		return err
	}

	client.messageID++

	return nil
}

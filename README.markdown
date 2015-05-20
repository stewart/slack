# slack

A small wrapper for Slack's [Real Time Messaging API][rtm].

It's pretty basic at the moment, patched together for personal use - if you'd like to see a feature added, let me know!

To use this, you'll need a Slack API token.
You can issue one by creating a [new Bot User integration](https://my.slack.com/services/new/bot).

## Installation

    $ go get github.com/stewart/slack

## Usage

Here's a basic example - see [GoDoc][] for more details on possible incoming events.

[GoDoc]: https://godoc.org/github.com/stewart/slack

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/stewart/slack"
	"github.com/stewart/slack/events"
)

var SLACK_TOKEN = os.Getenv("SLACK_TOKEN")

func main() {
	client := slack.NewClient(SLACK_TOKEN)

	if err := client.Connect(); err != nil {
		log.Fatal(err)
	}

	client.Loop()

	for {
		select {
		case err := <-client.Errors:
			log.Fatal(err)
		case msg := <-client.Incoming:
			parse(client, msg)
		}
	}
}

func parse(client *slack.Client, msg interface{}) {
	switch msg := msg.(type) {
	case events.Hello:
		fmt.Println("Slack says hello!")

	case events.PresenceChange:
		fmt.Println("There was a presence change!", msg.Presence)

	case events.Message:
		channel := msg.Channel
		text := "This is a basic message response"
		if err := client.SendMessage(channel, text); err != nil {
			fmt.Println("An error occured while responding", err)
		}

	default:
		fmt.Println("event received", msg)
	}
}
```

## License

MIT. For more details, see the `LICENSE` file.

[rtm]: https://api.slack.com/rtm

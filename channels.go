package slack

import "github.com/stewart/slack/types"

func listChannels(token string) ([]types.Channel, error) {
	resp, err := call("channels.list", token, map[string]string{})
	if err != nil {
		return []types.Channel{}, err
	}

	return resp.Channels, nil
}

package slack

import (
	"encoding/json"
	"errors"

	"github.com/stewart/slack/types"
)

type apiResponse struct {
	OK    bool   `json:"ok"`
	URL   string `json:"url"`
	Error string `json:"error"`

	// im.open response
	Channel struct {
		ID string `json:"id"`
	} `json:"channel"`

	// collections
	Members  []types.User    `json:"members"`
	Channels []types.Channel `json:"channels"`
	IMs      []types.IM      `json:"ims"`
}

func call(action, token string, params map[string]string) (*apiResponse, error) {
	resp := &apiResponse{}

	body, err := get(action, token, params)
	if err != nil {
		return resp, err
	}

	json.Unmarshal(body, &resp)

	if !resp.OK {
		return resp, errors.New("slack error: " + resp.Error)
	}

	return resp, nil
}

func startRtm(token string) (string, error) {
	resp, err := call("rtm.start", token, map[string]string{})
	if err != nil {
		return "", err
	}

	return resp.URL, nil
}

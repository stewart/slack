package slack

import "github.com/stewart/slack/types"

func listIms(token string) ([]types.IM, error) {
	resp, err := call("im.list", token, map[string]string{})
	if err != nil {
		return []types.IM{}, err
	}

	return resp.IMs, nil
}

func openIm(token, user string) (string, error) {
	resp, err := call("im.open", token, map[string]string{"user": user})
	if err != nil {
		return "", err
	}

	return resp.Channel.ID, nil
}

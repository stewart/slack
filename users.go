package slack

import "github.com/stewart/slack/types"

func listUsers(token string) ([]types.User, error) {
	resp, err := call("users.list", token, map[string]string{})
	if err != nil {
		return []types.User{}, err
	}

	return resp.Members, nil
}

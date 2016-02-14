package slack

import (
	"encoding/json"
	"errors"
)

type apiResponse struct {
	OK    bool   `json:"ok"`
	URL   string `json:"url"`
	Error string `json:"error"`
}

func requestRTM(token string) (string, error) {
	body, err := get("rtm.start", token, map[string]string{})
	if err != nil {
		return "", err
	}

	resp := apiResponse{}
	json.Unmarshal(body, &resp)

	if !resp.OK {
		return "", errors.New("slack error: " + resp.Error)
	}

	return resp.URL, nil
}

package slack

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const endpoint = "https://slack.com/api/"

type apiResponse struct {
	OK    bool   `json:"ok"`
	URL   string `json:"url"`
	Error string `json:"error"`
}

func requestRTM(token string) (string, error) {
	route, err := url.Parse(endpoint + "rtm.start")
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("token", token)
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

	resp := apiResponse{}
	json.Unmarshal(body, &resp)

	if !resp.OK {
		return "", errors.New("slack error: " + resp.Error)
	}

	return resp.URL, nil
}

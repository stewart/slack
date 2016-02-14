package slack

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const endpoint = "https://slack.com/api/"

func get(action, token string, params map[string]string) ([]byte, error) {
	route, err := url.Parse(endpoint + action)
	if err != nil {
		return []byte{}, err
	}

	query := url.Values{}

	query.Add("token", token)

	for key, value := range params {
		query.Add(key, value)
	}

	route.RawQuery = query.Encode()

	response, err := http.Get(route.String())
	if err != nil {
		return []byte{}, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

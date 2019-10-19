package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpConfig struct {
	URL string
}

func (config HttpConfig) GetUser() (UserInfo, error) {
	resp, err := makeRequest(http.MethodGet, config.URL, nil)
	if err != nil {
		return UserInfo{}, err
	}

	var userInfo UserInfo

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return UserInfo{}, fmt.Errorf("error decoding response %v", err)
	}

	return userInfo, nil
}

func makeRequest(method string, URL string, body io.Reader) (*http.Response, error) {

	request, err := http.NewRequest(method, URL, body)

	if err != nil {
		return nil, fmt.Errorf("error creating new HTTP request %v", err)
	}

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, fmt.Errorf("error getting response from service %v", err)
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("not found")
	}

	return resp, nil
}

package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thamaraiselvam/git-api-cli/cmd/types"
)

const githubHost = "https://api.github.com"

//Client is interface of service
type Client interface {
	GetUser() (types.UserInfo, error)
	GetGists() (types.Gists, error)
	GetFollowing() (types.FollowingUsers, error)
	GetFollowers() (types.Followers, error)
}

type config struct {
	BaseURL string
	URL     string
}

//CreateClient for making request
func CreateClient(path string) Client {
	return config{
		URL: githubHost + path,
	}
}

//GetFollowers fetches followers list of user
func (config config) GetFollowers() (types.Followers, error) {
	resp, err := makeRequest(http.MethodGet, config.URL, nil)

	if err != nil {
		return types.Followers{}, err
	}

	var followers types.Followers

	if err := json.NewDecoder(resp.Body).Decode(&followers); err != nil {
		return types.Followers{}, err
	}

	return followers, nil
}

//GetUser fetches user information from github.com
func (config config) GetUser() (types.UserInfo, error) {
	resp, err := makeRequest(http.MethodGet, config.URL, nil)
	if err != nil {
		return types.UserInfo{}, err
	}

	var userInfo types.UserInfo

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return types.UserInfo{}, fmt.Errorf("error decoding response %v", err)
	}

	return userInfo, nil
}

func (config config) GetGists() (types.Gists, error) {
	resp, err := makeRequest(http.MethodGet, config.URL, nil)
	if err != nil {
		return types.Gists{}, err
	}

	var gists types.Gists
	if err := json.NewDecoder(resp.Body).Decode(&gists); err != nil {
		return types.Gists{}, fmt.Errorf("error decoding response %v", err)
	}
	return gists, nil
}

//GetFollowing get following user information from github.com
func (config config) GetFollowing() (types.FollowingUsers, error) {
	resp, err := makeRequest(http.MethodGet, config.URL, nil)
	if err != nil {
		return types.FollowingUsers{}, err
	}
	userInfoList := make(types.FollowingUsers, 0)
	if err := json.NewDecoder(resp.Body).Decode(&userInfoList); err != nil {
		return types.FollowingUsers{}, fmt.Errorf("error decoding response %v", err)
	}

	return userInfoList, nil
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

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	return resp, nil
}

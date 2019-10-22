package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/thamaraiselvam/git-api-cli/cmd/types"
	"gopkg.in/h2non/gock.v1"
	"io"
	"net/http"
	"testing"
)

const githubHost = "https://api.github.com"

func TestHTTPConfig_GetUser(t *testing.T) {
	t.Run("should return user information on valid request", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/username").
			Reply(200).
			BodyString(`{"name": "username", "location": "location", "public_repos": 0}`)

		expectedUserInfo := types.UserInfo{
			Name:        "username",
			Location:    "location",
			PublicRepos: 0,
		}

		client := CreateClient("/users/username")

		actualUserInfo, err := client.GetUser()

		assert.NoError(t, err)
		assert.Equal(t, expectedUserInfo, actualUserInfo)
	})

	t.Run("should return not found on invalid username", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/username").
			Reply(404).
			BodyString(`{}`)

		client := CreateClient("/users/username")
		_, err := client.GetUser()

		assert.Error(t, err)
		assert.Equal(t, "404 Not Found", err.Error())
	})

	t.Run("should return error if response is not a valid json", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/username").
			Reply(200).
			BodyString(`string`)

		client := CreateClient("/users/username")
		_, err := client.GetUser()

		assert.Error(t, err)
		assert.Equal(t, "error decoding response invalid character 's' looking for beginning of value", err.Error())
	})

	t.Run("should return error if status is not 200", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/username").
			Reply(403).
			BodyString(`string`)

		client := CreateClient("/users/username")
		_, err := client.GetUser()

		assert.Error(t, err)
		assert.Equal(t, "403 Forbidden", err.Error())
	})

}

func TestConfig_GetFollowers(t *testing.T) {
	t.Run("Should return valid followers list on valid request", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/username/followers").
			Reply(200).
			BodyString(`[{"login": "username", "html_url": "https://github.com/username"}]`)

		expectedFollowers := types.Followers{
			{
				Name:    "username",
				HTMLURL: "https://github.com/username",
			},
		}

		client := CreateClient("/users/username/followers")
		actualFollowers, err := client.GetFollowers()

		assert.NoError(t, err)
		assert.Equal(t, expectedFollowers, actualFollowers)
	})

	t.Run("should return error and no followers when response is not a valid json", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/username/followers").
			Reply(200).
			BodyString(`invalid json`)

		client := CreateClient("/users/username/followers")
		followers, err := client.GetFollowers()

		assert.Error(t, err)
		assert.Equal(t, "invalid character 'i' looking for beginning of value", err.Error())
		assert.Equal(t, types.Followers{}, followers)
	})
}

func Test_makeRequest(t *testing.T) {

	gock.New(githubHost).
		Get("/users").
		Reply(200).
		BodyString(``)

	type args struct {
		method string
		URL    string
		body   io.Reader
	}
	tests := []struct {
		name        string
		args        args
		expected    http.Response
		expectedErr bool
	}{
		{
			name: "should pass on valid request",
			args: args{
				method: http.MethodGet,
				URL:    "https://api.github.com/users",
				body:   nil,
			},
			expected: http.Response{
				StatusCode: 200,
			},
			expectedErr: false,
		},
		{
			name: "should fail on invalid method type",
			args: args{
				method: "method",
				URL:    "https://api.github.com/users",
				body:   nil,
			},
			expected:    http.Response{},
			expectedErr: true,
		},
		{
			name: "should fail on invalid characters in url",
			args: args{
				method: http.MethodGet,
				URL:    "",
				body:   nil,
			},
			expected:    http.Response{},
			expectedErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := makeRequest(test.args.method, test.args.URL, test.args.body)

			if test.expectedErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			if actual != nil {
				assert.Equal(t, test.expected.StatusCode, actual.StatusCode)
			}
		})
	}
}

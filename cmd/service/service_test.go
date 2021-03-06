package service

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thamaraiselvam/git-api-cli/cmd/types"
	"gopkg.in/h2non/gock.v1"
)

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

func TestHTTPConfig_GetFollowing(t *testing.T) {
	t.Run("should return FollowingUsers on valid request", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/username/following").
			Reply(200).
			BodyString(`[{"login": "following", "html_url": "https://github.com/following"}]`)

		expectedUserInfo := types.FollowingUsers{{
			Name: "following",
			URL:  "https://github.com/following",
		}}

		client := CreateClient("/users/username/following")

		actualFollowingUserList, err := client.GetFollowing()

		assert.NoError(t, err)
		assert.Equal(t, expectedUserInfo, actualFollowingUserList)
	})

	t.Run("should return not found on invalid username", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/username/following").
			Reply(404).
			BodyString(`{}`)

		client := CreateClient("/users/username/following")
		_, err := client.GetFollowing()

		assert.Error(t, err)
		assert.Equal(t, "404 Not Found", err.Error())
	})

	t.Run("should return error if response is not a valid json", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/username/following").
			Reply(200).
			BodyString(`string`)

		client := CreateClient("/users/username/following")
		_, err := client.GetFollowing()

		assert.Error(t, err)
		assert.Equal(t, "error decoding response invalid character 's' looking for beginning of value", err.Error())
	})

}

func TestHTTPConfig_GetGists(t *testing.T) {
	t.Run("should return array of gists on valid request", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/username/gists").
			Reply(200).
			BodyString(`[{"description": "this is test gist", "created_at": "2019-10-22T14:29:31Z", "html_url": "http://github.com/username/test", "files":{"test.go":{"type": "go"}}}, {"description": "this is test gist2", "created_at": "2019-10-21T14:29:31Z", "html_url": "http://github.com/username/test2", "files":{"test1.go":{"type": "go"}, "test2.go":{"type": "go"}}}]`)

		expectedGists := types.Gists{
			{
				Files:       map[string]interface{}{"test.go": map[string]interface{}{"type": "go"}},
				URL:         "http://github.com/username/test",
				CreatedAt:   "2019-10-22T14:29:31Z",
				Description: "this is test gist",
			},
			{
				Files:       map[string]interface{}{"test1.go": map[string]interface{}{"type": "go"}, "test2.go": map[string]interface{}{"type": "go"}},
				URL:         "http://github.com/username/test2",
				CreatedAt:   "2019-10-21T14:29:31Z",
				Description: "this is test gist2",
			},
		}

		client := CreateClient("/users/username/gists")

		actualGists, err := client.GetGists()

		assert.NoError(t, err)
		assert.Equal(t, expectedGists, actualGists)
	})

	t.Run("should return error if response is not a valid json", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/username/gists").
			Reply(200).
			BodyString(`random`)

		client := CreateClient("/users/username/gists")
		_, err := client.GetGists()

		assert.Error(t, err)
		assert.Equal(t, "error decoding response invalid character 'r' looking for beginning of value", err.Error())
	})
	t.Run("should return not found on invalid username", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/username/gists").
			Reply(404).
			BodyString(`{}`)

		client := CreateClient("/users/username/gists")
		_, err := client.GetGists()

		assert.Error(t, err)
		assert.Equal(t, "404 Not Found", err.Error())
	})
}
func TestConfig_GetPRList(t *testing.T) {
	t.Run("Should return valid pull-request list on valid request", func(t *testing.T) {
		gock.New(githubHost).
			Get("/search/issues").
			Reply(200).
			BodyString(`{"items":[{ "title":"Test Issue","state":"open","pull_request":{"html_url":"www.github.com"}}]}`)

		client := CreateClient("/search/issues")
		actualPRList, err := client.GetPRList()
		assert.NoError(t, err)
		assert.Equal(t, "Test Issue", actualPRList.Items[0].Title)
		assert.Equal(t, "open", actualPRList.Items[0].State)
		assert.Equal(t, "www.github.com", actualPRList.Items[0].PullRequest.URL)
	})

	t.Run("should return error if response is not a valid json", func(t *testing.T) {
		gock.New(githubHost).
			Get("/search/issues").
			Reply(200).
			BodyString(`yo`)

		client := CreateClient("/search/issues")
		_, err := client.GetPRList()

		assert.Error(t, err)
		assert.Equal(t, "invalid character 'y' looking for beginning of value", err.Error())
	})
	t.Run("should return not found on invalid username", func(t *testing.T) {
		gock.New(githubHost).
			Get("/search/issues").
			Reply(404).
			BodyString(`{}`)

		client := CreateClient("/search/issues")
		_, err := client.GetPRList()

		assert.Error(t, err)
		assert.Equal(t, "404 Not Found", err.Error())
	})
}

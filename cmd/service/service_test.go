package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thamaraiselvam/git-api-cli/cmd/types"
	"gopkg.in/h2non/gock.v1"
)

const githubHost = "https://api.github.com"

func TestHTTPConfig_GetUser(t *testing.T) {
	t.Run("should return UserInfo on valid request", func(t *testing.T) {
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
		assert.Equal(t, "user not found", err.Error())
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

}

func TestHTTPConfig_GetPublicGists(t *testing.T) {
	t.Run("should return array of public gists on valid request", func(t *testing.T) {
		gock.New(githubHost).
			Get("/gists/public").
			Reply(200).
			BodyString(`[{"owner": {"login":"githubname"}, "description": "this is test gist", "created_at": "2019-10-22T14:29:31Z", "html_url": "http://github.com/user/test"}, {"owner": {"login":"githubname2" }, "description": "this is test gist2", "created_at": "2019-09-22T14:29:31Z", "html_url": "http://github.com/user/test2"}]`)

		expectedGists := []types.PublicGist{
			{
				Owner:       map[string]interface{}{"login": "githubname"},
				URL:         "http://github.com/user/test",
				CreatedAt:   "2019-10-22T14:29:31Z",
				Description: "this is test gist",
			},
			{
				Owner:       map[string]interface{}{"login": "githubname2"},
				URL:         "http://github.com/user/test2",
				CreatedAt:   "2019-09-22T14:29:31Z",
				Description: "this is test gist2",
			},
		}

		client := CreateClient("/gists/public")

		actualGists, err := client.GetPublicGists()

		assert.NoError(t, err)
		assert.Equal(t, expectedGists, actualGists)
	})

	t.Run("should return error if response is not a valid json", func(t *testing.T) {
		gock.New(githubHost).
			Get("/gists/public").
			Reply(200).
			BodyString(`random`)

		client := CreateClient("/gists/public")
		_, err := client.GetPublicGists()

		assert.Error(t, err)
		assert.Equal(t, "error decoding response invalid character 'r' looking for beginning of value", err.Error())
	})

}

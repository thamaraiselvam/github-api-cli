package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/thamaraiselvam/git-api-cli/cmd/types"
	"gopkg.in/h2non/gock.v1"
	"testing"
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

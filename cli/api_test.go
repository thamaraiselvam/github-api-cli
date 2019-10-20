package cli_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/thamaraiselvam/git-api-cli/cli"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

const githubHost = "https://api.github.com"

func TestHttpConfig_GetUser(t *testing.T) {
	t.Run("should return UserInfo on valid request", func(t *testing.T) {
		gock.New(githubHost).
			Get("/users/thamaraiselvam").
			Reply(200).
			BodyString(`{"name": "Thamaraiselvam", "location": "chennai", "public_repos": 90}`)

		expectedUserInfo := cli.UserInfo{
			Name:        "Thamaraiselvam",
			Location:    "chennai",
			PublicRepos: 90,
		}

		client := cli.HTTPConfig{
			URL: githubHost + "/users/thamaraiselvam",
		}

		actualUserInfo, err := client.GetUser()

		assert.NoError(t, err)
		assert.Equal(t, expectedUserInfo, actualUserInfo)
	})
}

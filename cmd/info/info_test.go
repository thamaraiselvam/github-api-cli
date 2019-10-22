package info

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/thamaraiselvam/git-api-cli/cmd/service"
	"github.com/thamaraiselvam/git-api-cli/cmd/types"
	mockClient "github.com/thamaraiselvam/git-api-cli/internal/mock/service"
	"testing"
)

func Test_getUserInfo(t *testing.T) {
	t.Run("should return userInfo on valid request", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedUserInfo := types.UserInfo{
			Name:        "name",
			Location:    "location",
			PublicRepos: 100,
		}

		client := mockClient.NewMockClient(ctrl)

		client.EXPECT().GetUser().Return(expectedUserInfo, nil)
		actualResult, err := getUserInfo(client)

		assert.NoError(t, err)
		assert.Equal(t, expectedUserInfo, actualResult)
	})

	t.Run("should return error on invalid request", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mockClient.NewMockClient(ctrl)

		client.EXPECT().GetUser().Return(types.UserInfo{}, fmt.Errorf("error"))
		_, err := getUserInfo(client)

		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())

	})
}

func Test_createClient(t *testing.T) {
	t.Run("should return a valid client", func(t *testing.T) {
		expectedClient := service.CreateClient("/users/username")
		actualClient := createClient("username")
		assert.Equal(t, expectedClient, actualClient)
	})
}

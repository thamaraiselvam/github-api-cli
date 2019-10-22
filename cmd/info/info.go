package info

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/thamaraiselvam/git-api-cli/cmd/service"
	"github.com/thamaraiselvam/git-api-cli/cmd/types"
	"os"
)

//Command fetch user data
func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Retrieve user profile information",
		Long:  "Retrieve user profile information such as name, location, public repos, followers, following",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("enter your name")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			client := service.CreateClient(fmt.Sprintf("/users/%s", args[0]))
			userInfo, err := getUserInfo(client)

			if err != nil {
				_ = fmt.Errorf("%v", err)
				os.Exit(1)
			}

			displayInfo(userInfo)
		},
	}
}

func getUserInfo(client service.Client) (types.UserInfo, error) {
	userInfo, err := client.GetUser()

	if err != nil {
		return types.UserInfo{}, err
	}

	return userInfo, err
}

func displayInfo(user types.UserInfo) {
	fmt.Println(user.Name)
	fmt.Println(user.Location)
	fmt.Println(user.PublicRepos)
}

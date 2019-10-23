package list

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/thamaraiselvam/git-api-cli/cmd/service"
	"github.com/thamaraiselvam/git-api-cli/cmd/types"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of ",
	Long:  "Retrieve user profile information such as name, location, public repos, followers, following",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("enter your name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client := createClient(args[0], args[1])
		repoInfo, err := getRepoInfo(client)

		if err != nil {
			_ = fmt.Errorf("%v", err)
			os.Exit(1)
		}

		displayInfo(repoInfo)
	},
}

//Command add list commands to root command
func Command() *cobra.Command {
	return listCmd
}

func createClient(name string, operation string) service.Client {
	return service.CreateClient(fmt.Sprintf("/users/%s/%s", name, operation))
}

func getRepoInfo(client service.Client) ([]types.RepoInfo, error) {
	repoInfo, err := client.GetRepos()

	if err != nil {
		return []types.RepoInfo{}, err
	}

	return repoInfo, err
}

func displayInfo(info []types.RepoInfo) {
	fmt.Println(info)
}

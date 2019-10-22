package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thamaraiselvam/git-api-cli/cmd/service"
)

func infoCmd() *cobra.Command {
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

			getRepoList(args[0])
		},
	}
}
func publicReposCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list gist",
		Short: "Retrieve public gists",
		Long:  "Retrieve public gists with owner github id, description, created date and url",
		Args:  nil,
		Run: func(cmd *cobra.Command, args []string) {
			getPublicGistList()
		},
	}
}

func getRepoList(name string) {
	client := service.CreateClient(fmt.Sprintf("/users/%s", name))
	userInfo, err := client.GetUser()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		os.Exit(1)
	}

	fmt.Println(userInfo.Name)
	fmt.Println(userInfo.Location)
	fmt.Println(userInfo.PublicRepos)
}

func getPublicGistList() {
	client := service.CreateClient("/gists/public")
	gists, err := client.GetPublicGists()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		os.Exit(1)
	}
	d := "Owner | Description | Created At | URL \n ---|----|-----\n"
	for _, gist := range gists {
		d += fmt.Sprintf("%s | %s | %s | %s", gist.Owner["login"].(string), gist.Description, gist.CreatedAt, gist.URL)
	}
	fmt.Println(d)
}

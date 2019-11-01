package list

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thamaraiselvam/git-api-cli/cmd/service"
)

func init() {
	listCmd.AddCommand(gistCmd)
}

var gistCmd = &cobra.Command{
	Use:   "list gists",
	Short: "Retrieve public gists",
	Long:  "Retrieve public gists with owner github id, description, created date and url",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("enter the username")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		getPublicGistList(args[0])
	},
}

func getPublicGistList(un string) {
	url := "/users/" + un + "/gists"
	client := service.CreateClient(url)
	gists, err := client.GetGists()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		os.Exit(1)
	}
	d := "URL | Description | Created At | Files \n -|-|-|-\n"
	for _, gist := range gists {
		d += fmt.Sprintf("%s | %s | %s | %d", gist.URL, gist.Description, gist.CreatedAt, len(gist.Files))
	}
	fmt.Println(d)
}

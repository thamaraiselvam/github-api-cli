package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thamaraiselvam/git-api-cli/cmd/service"
)

func publicGistsCmd() *cobra.Command {
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

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/thamaraiselvam/git-api-cli/cmd/service"
	"os"
	"text/tabwriter"
)

func followingCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "following <user name>",
		Short: "Retrieve list of following users",
		Long:  "Retrieve user profile information such as name and github url of following users",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("enter your name")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			getFollowingList(args[0])
		},
	}
}

func getFollowingList(name string) {
	client := service.CreateClient(fmt.Sprintf("/users/%s/following", name))
	userInfoList, err := client.GetFollowing()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		os.Exit(1)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	_, _ = fmt.Fprintln(w, "Name\tGithub Link")
	for _, userInfo := range userInfoList {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("%s\t%s", userInfo.Name, userInfo.URL))
	}
	_ = w.Flush()
}

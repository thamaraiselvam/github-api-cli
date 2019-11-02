package list

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/thamaraiselvam/git-api-cli/cmd/service"
	"github.com/thamaraiselvam/git-api-cli/cmd/types"
	"os"
	"text/tabwriter"
)

func init() {
	listCmd.AddCommand(followingCmd)
}

var followingCmd = &cobra.Command{
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

func getFollowingList(name string) {
	client := service.CreateClient(fmt.Sprintf("/users/%s/following", name))
	users, err := client.GetFollowing()
	if err != nil {
		_ = fmt.Errorf("%v", err)
		os.Exit(1)
	}
	displayFollowingList(users)
}

func displayFollowingList(users types.FollowingUsers) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	_, _ = fmt.Fprintln(w, "Name\tGithub Link")
	for _, user := range users {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("%s\t%s", user.Name, user.URL))
	}
	_ = w.Flush()
}

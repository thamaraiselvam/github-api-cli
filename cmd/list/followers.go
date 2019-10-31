package list

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/thamaraiselvam/git-api-cli/cmd/service"
	"github.com/thamaraiselvam/git-api-cli/cmd/types"
	"text/tabwriter"
	"os"
)

func init() {
	listCmd.AddCommand(followersCmd)
}

var followersCmd = &cobra.Command{
	Use:   "followers",
	Short: "list followers of current user",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("enter your name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client := service.CreateClient(fmt.Sprintf("/users/%s/followers", args[0]))
		followers, err := client.GetFollowers()
		if err != nil {
			_ = fmt.Errorf("%v", err)
			os.Exit(1)
		}

		displayFollowers(followers)
	},
}

func displayFollowers(followers types.Followers) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	_, _ = fmt.Fprintln(w, "Name\tGithub Link")
	for _, follower := range followers {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("%s\t%s", follower.Name, follower.HTMLURL))
	}
	_ = w.Flush()
}

package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

//FollowingUser contains following user information
type FollowingUser struct {
	Name string `json:"login"`
	URL  string `json:"html_url"`
}

//FollowingUserList stores a list of FollowingUser
type FollowingUserList = []FollowingUser

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
			config := createConfig()
			getFollowingList(config, args[0])
		},
	}
}

func getFollowingList(httpConfig HTTPConfig, name string) {
	httpConfig.URL = fmt.Sprintf("%s/users/%s/following", httpConfig.BaseURL, name)
	userInfoList, err := httpConfig.GetFollowing()
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

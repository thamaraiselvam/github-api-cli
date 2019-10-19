package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

//UserInfo contains all user related information
type UserInfo struct {
	Name string `json:"name"`
	Location string `json:"location"`
	PublicRepos int `json:"public_repos"`
}

var versionCmd = &cobra.Command{
	Use:   "user",
	Short: "Retrieve user profile information",
	Long: "Retrieve user profile information such as name, city, public repos, followers, following",
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

func getRepoList(name string){

	client := HttpConfig{URL:  fmt.Sprintf("https://api.github.com/users/%s", name)}
	userInfo, err := client.GetUser()
	if err != nil{
		_ = fmt.Errorf("%v", err)
		os.Exit(1)
	}

	fmt.Println(userInfo.Name)
	fmt.Println(userInfo.Location)
	fmt.Println(userInfo.PublicRepos)
}


package list

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of ",
	Long:  "Retrieve user profile information such as name, location, public repos, followers, following",
}

//Command add list commands to root command
func Command() *cobra.Command {
	return listCmd
}

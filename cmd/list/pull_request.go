package list

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thamaraiselvam/git-api-cli/cmd/service"
	"github.com/thamaraiselvam/git-api-cli/cmd/types"
	"github.com/thamaraiselvam/git-api-cli/cmd/util"
)

func init() {
	listCmd.AddCommand(prInfoCmd)
}

var prInfoCmd = &cobra.Command{
	Use:   "pr",
	Short: "list pull-requests of current user",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("enter your name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client := service.CreateClient(fmt.Sprintf("/search/issues?q=type:pr+author:%%22%s%%22&sort=created&order=desc", args[0]))
		prList, err := client.GetPRList()
		if err != nil {
			_ = fmt.Errorf("%v", err)
			os.Exit(1)
		}

		displayPRList(prList)
	},
}

func displayPRList(prList types.PRItemList) {

	table := util.CreateTable()
	table.SetHeader([]string{"Title", "URL"})
	for i := 0; i < len(prList.Items); i++ {
		table.Append([]string{prList.Items[i].Title, prList.Items[i].PullRequest.URL})
	}
	table.Render()

}

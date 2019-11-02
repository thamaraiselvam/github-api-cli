package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thamaraiselvam/git-api-cli/cmd/info"
	"github.com/thamaraiselvam/git-api-cli/cmd/list"
)

var rootCmd = &cobra.Command{}

//Execute on rootCmd
func Execute() {
	commands := []func() *cobra.Command{info.Command, list.Command}

	for _, command := range commands {
		rootCmd.AddCommand(command())
	}

	if err := rootCmd.Execute(); err != nil {
		_ = fmt.Errorf("github-api-cli failed with following errors")
		_ = fmt.Errorf(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

//Execute on rootCmd
func Execute() {
	commands := []func() *cobra.Command{infoCmd, publicGistsCmd}

	for _, command := range commands {
		rootCmd.AddCommand(command())
	}

	if err := rootCmd.Execute(); err != nil {
		_ = fmt.Errorf("github-api-cli failed with following errors")
		_ = fmt.Errorf(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
}

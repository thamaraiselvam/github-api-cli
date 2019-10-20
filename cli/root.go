package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{}

//Execute on rootCmd
func Execute() {
	commands := []func() *cobra.Command{infoCmd}

	for _, command := range commands {
		rootCmd.AddCommand(command())
	}

	if err := rootCmd.Execute(); err != nil {
		_ = fmt.Errorf("github-api-cli failed with following errors")
		_ = fmt.Errorf(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
}

func createConfig() HTTPConfig {
	return HTTPConfig{BaseURL: fmt.Sprintf("https://api.github.com")}
}

package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{}

//Execute on root cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_ = fmt.Errorf("gitapi failed with following errors")
		_ = fmt.Errorf(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
}

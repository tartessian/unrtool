package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of unrtool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tartessian's unrtool v0.1.0")
	},
}

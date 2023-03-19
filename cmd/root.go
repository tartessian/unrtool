package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "unrtool",
	Short:   "Import and export StaticMeshes to/from a UNR Lineage 2 map file.",
	Version: "0.1.0",
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	// rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

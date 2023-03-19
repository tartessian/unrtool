package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "unrtool",
	Short:   "Import and export StaticMeshes to/from a UNR Lineage 2 map file",
	Version: "0.1.0",
}

func init() {
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Execute()
}

package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tartessian/unrtool/internal/crypt"
)

var exportCmd = &cobra.Command{
	Use:   "export [file]",
	Short: "Export the StaticMeshes contained on a UNR file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		filename := filepath.Base(path)
		is, err := crypt.Decrypt(file, filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		destFile := "dec_" + filename
		newfile, err := os.Create(destFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer newfile.Close()

		buffer := make([]byte, 0x1000)
		for {
			r, err := is.Read(buffer)
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			_, err = newfile.Write(buffer[:r])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	},
}

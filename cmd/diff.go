package cmd

import (
	"github.com/jmendick/echodiff/internal/diff"
	"fmt"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff <file1> <file2>",
	Short: "Show differences between two files",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		file1, file2 := args[0], args[1]
		result, err := diff.CompareFiles(file1, file2)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)
}

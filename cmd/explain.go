package cmd

import (
	"github.com/jmendick/echodiff/internal/explain"
	"fmt"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain <file1> <file2>",
	Short: "Explain the changes between two files in natural language",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		file1, file2 := args[0], args[1]
		summary, err := explain.GenerateSummary(file1, file2)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(summary)
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
}

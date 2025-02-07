package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "echodiff",
	Short: "EchoDiff - A smarter diff tool with explanations",
	Long:  "EchoDiff is a CLI tool that not only shows diffs but explains changes in human-readable form.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'echodiff diff <file1> <file2>' to compare files.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

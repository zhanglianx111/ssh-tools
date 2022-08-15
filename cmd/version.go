package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Author    string
	GoVersion string
	CommitID  string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: string("打印版本信息"),
	Long:  string("\n打印版本信息"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Author: %s\n", Author)
		fmt.Printf("GoVersion: %s\n", GoVersion)
		fmt.Printf("CommitID: %s\n", CommitID)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

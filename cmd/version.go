package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: string("short: version information"),
	Long: string("long: Print version of ssh-tools"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

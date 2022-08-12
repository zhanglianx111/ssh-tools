package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use: string("update"),
	Short: string("short: ssh-tools update"),
	Long: string("long: ssh-tools update"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ssh-tools update")
	},
}
func init() {
	rootCmd.AddCommand(updateCmd)
	//updateCmd.Flags().String("")
}


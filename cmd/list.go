package cmd

import (
	"github.com/spf13/cobra"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "列出主机列表",
	Long:  "列出主机列表",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	rootCmd.AddCommand(cmdList)
}

func list() {
	machines := getAll()
	show(machines, "false")
}

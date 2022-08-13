package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use: string("update"),
	Short: string("更新数据库中到主机信息"),
	Long: string("\n更新数据库中到主机信息"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ssh-tools update")
	},
}
func init() {
	rootCmd.AddCommand(updateCmd)
	//updateCmd.Flags().String("")
}

// 支持更新用户名、密码和描述
func update() {

}
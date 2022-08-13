package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net"
	"os"
)

var addCmd = &cobra.Command{
	Use: string("add"),
	Short: string("将主机信息添加到数据库中"),
	Long: string("\n将主机信息添加到数据库中"),
	Run: func(cmd *cobra.Command, args []string) {
		add(cmd)
	},
}
func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("description", "d", "", "description about machine")

	addCmd.Flags().StringP("username", "u", "", "username for ssh")
	addCmd.MarkFlagRequired("username")

	addCmd.Flags().StringP("password", "p", "","password for ssh")
	addCmd.MarkFlagRequired("password")

	addCmd.Flags().StringP("ip", "i", "", "ip of machine")
	addCmd.MarkFlagRequired("ip")

}

type A struct {
	Name string
}
func add(cmd *cobra.Command) {
	ipAdd, err := cmd.Flags().GetString("ip")
	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	// TODO: check ip format
	ip := net.ParseIP(ipAdd)
	if ip == nil {
		fmt.Println("ip地址格式错误")
		os.Exit(1)
	}

	desc, err := cmd.Flags().GetString("description")
	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	username, err := cmd.Flags().GetString("username")
	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	passwd, err := cmd.Flags().GetString("password")
	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	db, _ := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	db.Create(&Machine{Ip: ip.String(), Description: desc, User: username, Password: passwd})
}

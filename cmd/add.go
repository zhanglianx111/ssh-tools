package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var addCmd = &cobra.Command{
	Use: string("add"),
	Short: string("short: ssh-tools add xxxx"),
	Long: string("long: ssh-tools add xxx"),
	Run: func(cmd *cobra.Command, args []string) {
		add(cmd)
	},
}
func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("description", "d", "", "description about machine")
	addCmd.Flags().StringP("username", "u", "", "username for ssh")
	addCmd.Flags().StringP("password", "p", "","password for ssh")
	addCmd.Flags().StringP("ip", "i", "", "ip of machine")
}

type A struct {
	Name string
}
func add(cmd *cobra.Command) {
	ip, err := cmd.Flags().GetString("ip")
	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	if 0 == len(ip) {
		fmt.Println("请输入ip地址、用户名和密码")
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

	db.Create(&Machine{Ip: ip, Description: desc, User: username, Password: passwd})
	//db.Commit()
}

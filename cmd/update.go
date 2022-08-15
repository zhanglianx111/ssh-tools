package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var updateCmd = &cobra.Command{
	Use:   string("update"),
	Short: string("更新数据库中的主机信息"),
	Long:  string("\n更新数据库中的主机信息"),
	Run: func(cmd *cobra.Command, args []string) {
		update(cmd, args)
	},
}

var help = `
Usage:
  use: update ip [flags]

Flags:
  -d, --description string   description
  -h, --help                 help for update
  -p, --password string      password
  -u, --username string      username
`

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("username", "u", "", "username")
	updateCmd.Flags().StringP("password", "p", "", "password")
	updateCmd.Flags().StringP("description", "d", "", "description")
	updateCmd.SetHelpTemplate(help)
}

// 支持更新用户名、密码和描述
func update(cmd *cobra.Command, args []string) {
	if len(args) == 0 || len(args) != 1 {
		cmd.Help()
		os.Exit(0)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}

	ip := args[0]
	var m Machine
	/*
		if tx := db.First(&m, "ip = ?", ip); tx.Error.Error() != "" {
			//fmt.Println(tx.Error.Error())
			os.Exit(1)
		}

	*/

	user, err := cmd.Flags().GetString("username")
	if err != nil {
		fmt.Println(err.Error())
	}
	passwd, err := cmd.Flags().GetString("password")
	if err != nil {
		fmt.Println(err.Error())
	}
	desc, err := cmd.Flags().GetString("description")
	if err != nil {
		fmt.Println(err.Error())
	}

	db.Model(&m).Where("ip = ?", ip).Updates(Machine{Ip: ip, User: user, Password: passwd, Description: desc})
}

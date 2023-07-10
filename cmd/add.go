package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var addCmd = &cobra.Command{
	Use:   string("add"),
	Short: string("将主机信息添加到数据库中"),
	Long:  string("\n将主机信息添加到数据库中"),
	Run: func(cmd *cobra.Command, args []string) {
		add(cmd)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("description", "d", "", "description about machine")

	addCmd.Flags().StringP("username", "u", "", "username for ssh")
	addCmd.MarkFlagRequired("username")

	addCmd.Flags().StringP("password", "p", "", "password for ssh")
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
		log.Print(err.Error())
		return
	}

	// TODO: check ip format
	ip := net.ParseIP(ipAdd)
	if ip == nil {
		log.Println("ip地址格式错误")
		return
	}

	desc, err := cmd.Flags().GetString("description")
	if err != nil {
		log.Println(err.Error())
		return
	}

	username, err := cmd.Flags().GetString("username")
	if err != nil {
		log.Println(err.Error())
		return
	}

	passwd, err := cmd.Flags().GetString("password")
	if err != nil {
		log.Println(err.Error())
		return
	}

	db, _ := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	db.Create(&Machine{Ip: ip.String(), Description: desc, User: username, Password: passwd})
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

const createTable = `
	CREATE TABLE IF NOT EXISTS machines (
  	ip TEXT NOT NULL PRIMARY KEY,
	description TEXT,
  	user TEXT NOT NULL,
	password TEXT NOT NULL
  );
`
var initCmd = &cobra.Command{
	Use: string("init"),
	Short: string("初始化数据库"),
	Long: string("\n初始化数据库"),
	Run: func(cmd *cobra.Command, args []string) {
		initEnv()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initEnv() {
	_, err := os.Stat(dbPath)
	if err != nil && os.IsNotExist(err) {
		if _, errCreate := os.Create(dbPath); errCreate != nil {
			fmt.Printf("创建数据库失败\n")
			os.Exit(5)
		}
		db, _ := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		t := db.Exec(createTable)
		if t == nil {
			fmt.Println("创建表失败")
		}
	}
	fmt.Sprintln()
}
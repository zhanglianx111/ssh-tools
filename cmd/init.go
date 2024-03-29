package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const createTable = `
	CREATE TABLE IF NOT EXISTS machines (
  	ip TEXT NOT NULL PRIMARY KEY,
	description TEXT,
  	user TEXT NOT NULL,
	password TEXT NOT NULL,
	status TEXT
  );
`

var initCmd = &cobra.Command{
	Use:   string("init"),
	Short: string("初始化数据库"),
	Long:  string("\n初始化数据库"),
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
			log.Printf("创建数据库失败\n")
			return
		}
		db, _ := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		t := db.Exec(createTable)
		if t == nil {
			log.Println("创建表失败")
		}
	}
	fmt.Sprintln()
}

package cmd

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

type Machine struct {
	Ip 			string
	Description string
	User 		string
	Password 	string
}

const sshPort = "22"

func getAll() []Machine {
	//var name, desc, user, passwd, ip string
	db, err := gorm.Open(sqlite.Open(dbPath), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(7)
	}
	machines := []Machine{}
	db.Find(&machines)
	if len(machines) == 0 {
		fmt.Println("未找到主机")
		return machines
	}
	for i:=0; i<len(machines); i++ {
		fmt.Println(i, machines[i].Ip, machines[i].Description)
	}
	return machines
}
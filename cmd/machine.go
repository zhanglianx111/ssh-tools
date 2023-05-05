package cmd

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Machine struct {
	Ip          string
	Description string
	User        string
	Password    string
	Status      bool
}

const sshPort = "22"

func getAll() []Machine {
	machines := []Machine{}
	//var name, desc, user, passwd, ip string
	db, err := gorm.Open(sqlite.Open(dbPath), nil)
	if err != nil {
		log.Println(err.Error())
		return machines
	}

	db.Find(&machines)
	if len(machines) == 0 {
		fmt.Println("未找到主机")
	}

	return machines
}

func show(m []Machine, flag string) {
	if len(m) == 0 {
		return
	}

	if flag == "true" {
		fmt.Println("序号\tIP\t\t状态\t描述")
		for i := 0; i < len(m); i++ {
			fmt.Printf("%d ---> %s\t%t\t%s\n", i, m[i].Ip, m[i].Status, m[i].Description)
		}
	} else {
		fmt.Println("序号\tIP\t\t描述")
		for i := 0; i < len(m); i++ {
			fmt.Printf("%d ---> %s\t%s\n", i, m[i].Ip, m[i].Description)
		}
	}
	return
}

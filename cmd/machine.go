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
		fmt.Println("序号\tIP\t\t描述\t状态")
		for i := 0; i < len(m); i++ {
			fmt.Printf("%d ---> %s\t%s\t%t\n", i, m[i].Ip, m[i].Description, m[i].Status)
		}
	} else {
		fmt.Println("序号\tIP\t\t描述")
		for i := 0; i < len(m); i++ {
			fmt.Printf("%d ---> %s\t%s\n", i, m[i].Ip, m[i].Description)
		}
	}
	return
}

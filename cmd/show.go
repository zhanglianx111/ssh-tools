package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   string("show"),
	Short: string("显示主机用户名和密码信息"),
	Long:  string("\n显示主机用户名和密码信息"),
	Run: func(cmd *cobra.Command, args []string) {
		showOne()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func showOne() {
	machines := getAll()
	show(machines, false)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("请输入要显示主机信息的序号: ")

	next:
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err.Error())
			return
		}

		// show machine
		choose := strings.Split(text, "\n")[0]
		if 0 == len(choose) {
			return
		}

		index, err := strconv.ParseInt(choose, 10, 32)
		if err != nil {
			//log.Println(err.Error())
			fmt.Printf("序号选择错误，请重新选择：")
			goto next
		}

		if 0 <= int(index) && int(index) < len(machines) {
			fmt.Printf("user: %s, password: %s\n\n", machines[index].User, machines[index].Password)
		} else {
			fmt.Printf("序号选择错误，请重新选择：")
			goto next
		}
	}
}

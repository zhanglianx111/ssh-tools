package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)

var deleteCmd = &cobra.Command{
	Use: string("delete"),
	Short: string("删除数据库中的主机信息"),
	Long: string("\n删除数据库中的主机信息"),
	Run: func(cmd *cobra.Command, args []string) {
		delete()
	},
}
func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delete() {
	machines := getAll()
	show(machines)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("请输入要删除主机的序号: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(9)
		}

		// delete machine
		choose := strings.Split(text, "\n")[0]
		if 0 == len(choose) {
			os.Exit(0)
		}

		index, err := strconv.Atoi(choose)
		if err != nil {
			fmt.Println(err)
			os.Exit(10)
		}

		db, err := gorm.Open(sqlite.Open(dbPath), nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(7)
		}
		db.Where("ip = ?", machines[index].Ip).Delete(&Machine{})
	}
}
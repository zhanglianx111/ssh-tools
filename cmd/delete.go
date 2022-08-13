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
	Short: string("short: ssh-tools delete xxxx"),
	Long: string("long: ssh-tools delete xxx"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ssh-tools delete xxxx")
		delete()
	},
}
func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delete() {
	machines := getAll()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("输入: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(9)
		}

		// ssh login machine
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
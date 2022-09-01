package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	dbFile = ".ssh-machines.db"
	dbPath string
)

const (
	TerminalWidth  = 200
	TerminalHeight = 51
	SshDailTimeout = 5
)

var rootCmd = cobra.Command{
	Use:   string("use: ssh-tools"),
	Short: string("short: ssh-tools"),
	Long:  string("\nssh登陆主机工具"),

	//Run: func(cmd *cobra.Command, args []string) {
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err.Error())
		return
	}
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	dbPath = homeDir + "/" + dbFile
}

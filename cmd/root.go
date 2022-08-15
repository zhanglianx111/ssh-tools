package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	dbFile = ".ssh-machines.db"
	dbPath string
)

const (
	TerminalWidth  = 200
	TerminalHeight = 51
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
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		os.Exit(1)
	}
	dbPath = homeDir + "/" + dbFile
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var dbFile = ".ssh-machines.db"
var dbPath string

var rootCmd = cobra.Command{
	Use: string("use: ssh-tools"),
	Short: string("short: ssh-tools"),
	Long: string("long: ssh tools"),

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

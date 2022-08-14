package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"strconv"
	"strings"
)

var listCmd = &cobra.Command{
	Use: string("login"),
	Short: string("登陆主机"),
	Long: string("\n登陆主机"),
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}
func init() {
	rootCmd.AddCommand(listCmd)
}

func login() {
	machines := getAll()
	if len(machines) == 0 {
		return
	}
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
	doSsh(machines[index].Ip, machines[index].User, machines[index].Password)
}

func doSsh(ip, user, passwd string) {
	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer terminal.Restore(fd, oldState)

	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", ip + ":" + sshPort, &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}
	defer client.Close()

	// 建立新会话
	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,  // 禁用回显（0禁用，1启动）
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
		ssh.VSTATUS: 		1,
	}

	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		log.Fatalf(err.Error())
	}
	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	if err = session.RequestPty("xterm", termWidth, termHeight, modes); err != nil {
		fmt.Println("1", err.Error())
	}
	if err = session.Shell(); err != nil {
		fmt.Println("2", err.Error())
	}
	if err = session.Wait(); err != nil {
		fmt.Println("3", err.Error())
	}
}
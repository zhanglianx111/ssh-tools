package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
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
	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", ip + ":" + sshPort, &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	// 建立新会话
	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}

	session.Stdout = os.Stdout // 会话输出关联到系统标准输出设备
	session.Stderr = os.Stderr // 会话错误输出关联到系统标准错误输出设备
	session.Stdin = os.Stdin   // 会话输入关联到系统标准输入设备
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,  // 禁用回显（0禁用，1启动）
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
	}
	if err = session.RequestPty("linux", 32, 160, modes); err != nil {
		log.Fatalf("request pty error: %s", err.Error())
	}
	if err = session.Shell(); err != nil {
		log.Fatalf("start shell error: %s", err.Error())
	}
	if err = session.Wait(); err != nil {
		log.Fatalf("return error: %s", err.Error())
	}
}
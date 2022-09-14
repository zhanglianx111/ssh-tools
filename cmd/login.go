package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

var listCmd = &cobra.Command{
	Use:   string("login"),
	Short: string("登陆主机"),
	Long:  string("\n登陆主机"),
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

type window struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func login() {
	machines := getAll()
	show(machines)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入序号: ")
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
	client, err := ssh.Dial("tcp", ip+":"+sshPort, &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: 		SshDailTimeout * time.Second,
	})
	if err != nil {
		log.Printf("SSH dial error: %s", err.Error())
		return
	}
	defer client.Close()

	// 建立新会话
	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		log.Printf("new session error: %s", err.Error())
		return
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // 禁用回显（0禁用，1启动）
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
		ssh.VSTATUS:       1,
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	height, width, err := getLinesAndColumns()
	if err != nil {
		return
	}
	fmt.Println(height, width)
	if err = session.RequestPty("xterm", height, width, modes); err != nil {
		fmt.Println("1", err.Error())
		return
	}
	if err = session.Shell(); err != nil {
		fmt.Println("2", err.Error())
		return
	}

	session.Wait()
}
/*
func getLinesAndColumns() (int, int) {
	l := os.Getenv("LINES")
	c := os.Getenv("COLUMNS")
	fmt.Println("sss", l, c)
	lines, err := strconv.Atoi(l)
	if err != nil {
		lines = TerminalHeight
	}
	colcumns, err := strconv.Atoi(c)
	if err != nil {
		colcumns = TerminalWidth
	}
	return lines, colcumns
}
 */

func getLinesAndColumns() (int, int, error) {
	w := new(window)
	tio := syscall.TIOCGWINSZ
	if runtime.GOOS == "darwin" {
		tio = TIOCGWINSZ_OSX
	}
	res, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(tio),
		uintptr(unsafe.Pointer(w)),
	)
	if int(res) == -1 {
		return 0, 0, err
	}
	return int(w.Row), int(w.Col), nil
}
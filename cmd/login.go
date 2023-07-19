package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/spf13/cobra"
	"github.com/zhanglianx111/ssh-tools/utils"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var listCmd = &cobra.Command{
	Use:   string("login"),
	Short: string("登陆主机"),
	Long:  string("\n登陆主机"),
	Run: func(cmd *cobra.Command, args []string) {
		statusFlag, err := cmd.Flags().GetBool("ping")
		if err != nil {
			log.Print(err.Error())
			return
		}

		if len(args) == 1 && args[0] == "false" {
			statusFlag = false
		}
		login(statusFlag)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("ping", "p", true, "get status by pinging machine, default vaule: true")
}

type window struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func login(flag bool) {
	machines := getAll()

	// update machine connneted status
	if flag {
		machineIPs := []string{}
		for _, m := range machines {
			machineIPs = append(machineIPs, m.Ip)
		}

		status := utils.PingAll(machineIPs)

		for i := 0; i < len(machines); i++ {
			machines[i].Status = status[machines[i].Ip]
		}
	}
	show(machines, flag)

	fmt.Print("请输入要登陆主机的序号: ")

input:
	reader := bufio.NewReader(os.Stdin)
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
	if index < 0 || index >= len(machines) {
		fmt.Print("输入的主机序号错误，请重新输入: ")
		goto input
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
		Timeout:         SshDailTimeout * time.Second,
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

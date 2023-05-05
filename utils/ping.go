package utils

import (
	"net"
	"time"
)

func PingOne(machine string) bool {
	return pingOnce(machine)
}

func PingAll(machines []string) map[string]bool {
	status := make(map[string]bool, len(machines))
	for _, m := range machines {
		state := pingOnce(m)
		status[m] = state
	}
	return status
}

func pingOnce(machine string) bool {
	conn, err := net.DialTimeout("tcp", machine+":"+"22", 10*time.Millisecond)
	if err != nil {
		return false
	} else {
		if conn != nil {
			conn.Close()
			return true
		} else {
			return false
		}
	}
}

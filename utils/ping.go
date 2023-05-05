package utils

import (
	"fmt"

	"github.com/go-ping/ping"
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
	fmt.Println(status)
	return status
}

func pingOnce(machine string) (stat bool) {
	pinger, err := ping.NewPinger(machine)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics()
	stat = true
	fmt.Println(stats, stat)

	return stat
}

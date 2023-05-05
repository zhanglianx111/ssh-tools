package utils

import "testing"

func TestA(t *testing.T) {
	t.Log(PingOne("www.baidu.com"))
	machines := []string{"127.0.0.1", "www.baidu.com"}
	t.Log(PingAll(machines))
}

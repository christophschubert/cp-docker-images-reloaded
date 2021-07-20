package main

import (
	"fmt"
	"net"
	"time"
)

func waitForService(host string, port string, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Println("Connection could not be opened")
		return false
	} else {
		fmt.Println(conn.LocalAddr())
		conn.Close()
		return true
	}
}

func main() {
	waitForService("localhost", "9092", 5 * time.Second)
}

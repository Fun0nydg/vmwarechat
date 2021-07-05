package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start the server")
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Printf("error listen:%s", err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("the error is :%s", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}
func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v\n", string(buf[:len]))
	}
}

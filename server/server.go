package main

import (
	"fmt"
	"net"
	"strings"
)

var connmap map[string]net.Conn = make(map[string]net.Conn)

func main() {
	fmt.Println("start the server")
	listener, err := net.Listen("tcp", "127.0.0.1:50001")
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
		// who := conn.RemoteAddr()
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		msg_str := strings.Split(string(buf[:len]), "|")
		connmap[msg_str[0]] = conn
		for k, v := range connmap {
			fmt.Println(k, v)
			if k != msg_str[0] {
				v.Write([]byte("[" + msg_str[0] + "]:" + msg_str[1]))
			}
		}

		fmt.Println(msg_str)

	}
}

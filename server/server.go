package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var connmap map[string]net.Conn = make(map[string]net.Conn)

func main() {
	fmt.Println("start the server")
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
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
		fmt.Println("start doServerStuff")
		buf := make([]byte, 512)
		len, err := conn.Read(buf) //1
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}

		msg_str := strings.Split(string(buf[:len]), "|---|")
		fmt.Println(msg_str[1][:4])
		connmap[msg_str[0]] = conn
		fmt.Println(msg_str[0])
		switch msg_str[1][:5] {
		case "post ":
			for k, v := range connmap {
				if k != msg_str[0] {
					fmt.Println(k, v)
					v.Write([]byte("[" + msg_str[0] + "]:" + msg_str[1][4:]))
				}
			}
			break
		case "file ":
			fmt.Println("the file case in")
			buf := make([]byte, 4096)
			fmt.Println("start receive filename")
			len, err := conn.Read(buf) //2
			if err != nil {
				fmt.Println("the error is:", err.Error())
				break
			}
			filename := string(buf[:len])
			fmt.Println("the filename is :", filename)
			if filename != "" {
				fmt.Println("start send ok")
				// conn.Write([]byte("ok"))
				time.Sleep(1 * time.Second)
			}
			file, err := os.Create(filename)
			for {
				fmt.Println("start receive file")
				buf := make([]byte, 4096)
				len, err := conn.Read(buf)
				fmt.Println(string(buf[:len]))
				if string(buf[:len]) == "finish" {
					fmt.Println("the file is receive complete")
					break
				}
				if err != nil {
					fmt.Println("the error is:", err.Error())
					break
				}
				file.Write(buf[:len])
			}
			defer file.Close()
			fmt.Println("jump for")
			break
		}

		defer conn.Close()
		fmt.Println(msg_str)
	}
}

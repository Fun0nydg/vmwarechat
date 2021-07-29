package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var connmap map[string]net.Conn = make(map[string]net.Conn)

func sendfile(conn net.Conn, file *os.File) {
	fmt.Println("Start send file to client")
	_, err := io.Copy(conn, file)
	if err != nil {
		fmt.Println(err)
	}
}

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
		fmt.Println(conn.RemoteAddr(), "connect successed")
		go doServerStuff(conn)
	}
}
func doServerStuff(conn net.Conn) {
	for {
		fmt.Println("start service")
		buf := make([]byte, 512)
		len, err := conn.Read(buf) //1 conn
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		// fmt.Println(string(buf[:len]))
		msg_str := strings.Split(string(buf[:len]), "|---|")
		// fmt.Println(msg_str)
		connmap[msg_str[0]] = conn
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
			fmt.Println("server start receive file")
			serverftp, err := net.Listen("tcp", "0.0.0.0:8001")
			if err != nil {
				fmt.Println(err)
			}
			conn1, err := serverftp.Accept()
			if err != nil {
				fmt.Println(err)
			}
			filename := msg_str[1][5:]
			fmt.Println("receive filename:", filename)
			file, err := os.Create(filename)
			_, err = io.Copy(file, conn1)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("the file is receive complete")
			file.Close()
			serverftp.Close()
			break
		case "down ":
			fmt.Println("down case in")
			conn2, err := net.DialTimeout("tcp", conn.RemoteAddr().String(), 5*time.Second)
			if err != nil {
				fmt.Println(err)
				break
			}
			filename := msg_str[1][5:]
			file, err := os.Open(filename)
			if err != nil {
				fmt.Println(err)
				conn2.Write([]byte("fail"))
				break
			}
			fmt.Println("start send file:", filename)
			conn2.Write([]byte("success"))
			time.Sleep(1 * time.Second)
			sendfile(conn2, file)
			fmt.Println("end down case")
			file.Close()
			conn2.Close()
			break
		}
	}
}

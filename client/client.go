package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//设置参数
	ip := flag.String("ip", "127.0.0.1", "input ip address")
	port := flag.Uint("p", 50001, "port")
	strport := fmt.Sprint(*port)
	flag.Parse()
	fmt.Printf("now we connect to server:%s", *ip+":"+strport)
	conn, err := net.Dial("tcp", *ip+":"+strport)
	defer conn.Close()
	inputRd := bufio.NewReader(os.Stdin)
	fmt.Println("\ninput your name:")
	name, _ := inputRd.ReadString('\n')
	trimname := strings.Trim(name, "\n")
	fmt.Println("the name is :", trimname, "end")
	// conn.Write([]byte(trimname))
	go clientread(conn)
	fmt.Println("\nplease input your msg,enter Q is quit:")

	if err != nil {
		fmt.Printf("connect error msg:%s", err.Error())
	}

	for {
		input, err := inputRd.ReadString('\n')
		if err != nil {
			fmt.Println("error:", err.Error())
			break
		}
		triminput := strings.Trim(input, "\n")
		if triminput == "Q" {
			fmt.Println("quit")
			break
		}
		conn.Write([]byte(trimname + "|" + triminput))
	}
}

func clientread(conn net.Conn) {
	for {
		data := make([]byte, 512)
		redata, err := conn.Read(data)
		if redata == 0 || err != nil {
			break
		}
		fmt.Println(string(data[:redata]))
	}
}

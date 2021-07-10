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
	fmt.Println(trimname)
	conn.Write([]byte("name" + "|" + trimname))
	fmt.Println("\nplease input your msg,enter Q is quit:")

	if err != nil {
		fmt.Printf("connect error msg:%s", err.Error())
	}
	go clientread(conn)

	for {
		input, err := inputRd.ReadString('\n')
		if err != nil {
			fmt.Println("error:", err.Error())
			return
		}
		triminput := strings.Trim(input, "\n")
		if triminput == "Q" {
			fmt.Println("quit")
			return
		}
		conn.Write([]byte("say" + "|" + trimname + "|" + triminput))
	}
}

func clientread(conn net.Conn) {
	data := make([]byte, 512)
	redata, err := conn.Read(data)
	if redata == 0 || err != nil {
		return
	}
	fmt.Println(string(data[:redata]))
}

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
	ip := flag.String("ip", "", "")
	port := flag.Uint("p", 50001, "")
	strport := fmt.Sprint(*port)
	flag.Parse()
	fmt.Printf("now we connect to server:%s", *ip+":"+strport)
	conn, err := net.Dial("tcp", *ip+":"+strport)
	fmt.Println("\nplease input your msg,enter Q is quit:")
	if err != nil {
		fmt.Printf("connect error msg:%s", err.Error())
	}
	inputRd := bufio.NewReader(os.Stdin)
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
		conn.Write([]byte(triminput))
	}
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var input string

	conn, _ := net.Dial("tcp", "localhost:5000")
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

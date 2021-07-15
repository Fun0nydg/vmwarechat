package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func sendfile(conn net.Conn, filepath string) {
	fmt.Println("sendfile")
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("the error is :", err.Error())
	}
	buf := make([]byte, 4096)
	for {
		fmt.Println("client statrt to send")
		send, err := file.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("the file is transform complete!")
			time.Sleep(1 * time.Second)
			conn.Write([]byte("finish")) //4
			break
		}
		conn.Write(buf[:send]) //3
	}
}
func main() {
	//设置参数
	ip := flag.String("ip", "127.0.0.1", "input ip address")
	port := flag.Uint("p", 8000, "port")
	flag.Parse()
	strport := fmt.Sprint(*port)
	fmt.Println(*ip, strport)
	fmt.Printf("now we connect to server:%s", *ip+":"+strport)
	conn, err := net.DialTimeout("tcp", *ip+":"+strport, 5*time.Second)
	defer conn.Close()
	inputRd := bufio.NewReader(os.Stdin)
	fmt.Println("\ninput your name:")
	name, _ := inputRd.ReadString('\n')
	trimname := strings.Trim(name, "\n")
	fmt.Println("the name is :", trimname, "end")
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
		caseinput := triminput[:4]
		fmt.Println(caseinput)
		switch caseinput {
		case "post":
			conn.Write([]byte(trimname + "|---|" + triminput))
			// clientread(conn)
			fmt.Println("post end")
		case "quit":
			break
		case "file":
			conn.Write([]byte(trimname + "|---|" + triminput)) //1
			time.Sleep(1 * time.Second)
			filepath := triminput[5:]
			fmt.Println(filepath)
			fileinfo, err := os.Stat(filepath)
			if err != nil {
				fmt.Println("the error is :", err.Error())
				break
			}
			filename := fileinfo.Name()
			fmt.Println("start send filename")
			conn.Write([]byte(filename)) //2
			time.Sleep(1 * time.Second)
			if err != nil {
				fmt.Println("the error is", err.Error())
				break
			}
			fmt.Println("111")

			// buf := make([]byte, 512)
			// res, err := conn.Read(buf)
			// fmt.Println("receive", string(buf[:res]))
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			sendfile(conn, filename)
			// if string(buf[:res]) == "ok" {
			// 	fmt.Println("start send")
			// 	sendfile(conn, filename)
			// 	break
			// }

		}
		if triminput == "Q" {
			fmt.Println("quit")
			break
		}
		fmt.Println("switch end")

	}
}

func clientread(conn net.Conn) {
	for {
		data := make([]byte, 4096)
		redata, err := conn.Read(data)
		if redata == 0 || err != nil {
			break
		}
		fmt.Println(string(data[:redata]))
	}
}

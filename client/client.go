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

func sendfile(conn net.Conn, filepath string, filesize int64) {
	fmt.Println("sendfile")
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("the error is :", err.Error())
	}
	fmt.Println("the file size is:", filesize)
	buf := make([]byte, 100000)
	for {
		fmt.Println("client statrt to send")
		send, err := file.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("the file is transform complete!")
			time.Sleep(3 * time.Second)
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
	name = strings.Replace(name, "\r", "", -1)
	name = strings.Replace(name, "\n", "", -1)
	trimname := name
	if len(trimname) != 3 {
		fmt.Println("name length error,please input 3 length name,eg:aa1")
		return
	}
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
		// triminput := strings.Trim(input, "\n")
		input = strings.Replace(input, "\r", "", -1)
		input = strings.Replace(input, "\n", "", -1)
		triminput := input
		if len(triminput) <= 5 {
			fmt.Println("input error,try again,example: post 123 or file 123.txt")
			continue
		}
		caseinput := triminput[:5]
		if caseinput != "post " && caseinput != "file " {
			fmt.Println("error,please input post msg or file file.txt")
			continue
		}
		// fmt.Println(caseinput)
		switch caseinput {
		case "post ":
			conn.Write([]byte(trimname + "|---|" + triminput))
			// clientread(conn)
			fmt.Println("post end")
		case "quit ":
			break
		case "file ":

			filepath := triminput[5:]
			// fmt.Println(filepath)
			fileinfo, err := os.Stat(filepath)
			if err != nil {
				fmt.Println("the error is :", err.Error())
				break
			}
			filesize := fileinfo.Size()

			conn.Write([]byte(trimname + "|---|" + triminput)) //1
			time.Sleep(1 * time.Second)
			filename := fileinfo.Name()
			// fmt.Println("start send filename")

			if err != nil {
				fmt.Println("the error is", err.Error())
				break
			}
			conn.Write([]byte(filename)) //2
			time.Sleep(1 * time.Second)
			// fmt.Println("111")

			// buf := make([]byte, 512)
			// res, err := conn.Read(buf)
			// fmt.Println("receive", string(buf[:res]))
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			sendfile(conn, filename, filesize)
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
		// fmt.Println("switch end")

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

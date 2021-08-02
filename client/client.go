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

func sendfile(conn1 net.Conn, filename string) {
	fmt.Println("start sendfile")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("the error is :", err.Error())
	}
	_, err = io.Copy(conn1, file)
	file.Close()
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
	fmt.Println("\nplease input your msg,eg:post xxxx or file filename.txt")
	if err != nil {
		fmt.Printf("connect error msg:%s", err.Error())
	}

	for {
		fmt.Println("choose the method")
		input, err := inputRd.ReadString('\n')
		if err != nil {
			fmt.Println("error:", err.Error())
			break
		}
		input = strings.Replace(input, "\r", "", -1)
		input = strings.Replace(input, "\n", "", -1)
		triminput := input
		if len(triminput) <= 5 {
			fmt.Println("input error,try again,example: post 123 or file 123.txt")
			continue
		}
		caseinput := triminput[:5]
		if caseinput != "post " && caseinput != "file " && caseinput != "down " {
			fmt.Println("error,please input post msg or file file.txt")
			continue
		}
		switch caseinput {
		case "post ":
			conn.Write([]byte(trimname + "|---|" + triminput))
			fmt.Println("post end")
		case "quit ":
			break
		case "file ":
			fmt.Println("case file")
			filepath := triminput[5:]
			fileinfo, err := os.Stat(filepath)
			if err != nil {
				fmt.Println(err)
				break
			}
			conn.Write([]byte(trimname + "|---|" + triminput)) //1 conn
			time.Sleep(1 * time.Second)
			filename := fileinfo.Name()
			fmt.Println("the filename is:", filename)
			conn1, err := net.DialTimeout("tcp", *ip+":8001", 5*time.Second)
			if err != nil {
				fmt.Println(err)
			}
			sendfile(conn1, filepath)
			conn1.Close()
			break
		case "down ":
			fmt.Println("down case in")
			conn.Write([]byte(trimname + "|---|" + triminput))
			serverrec, err := net.Listen("tcp", "0.0.0.0:8002")
			if err != nil {
				fmt.Println(err)
			}
			conn2, err := serverrec.Accept()
			filename := triminput[5:]
			file, err := os.Create(filename)
			if err != nil {
				fmt.Println(err)
			}
			buf := make([]byte, 512)
			len, err := conn2.Read(buf)
			fmt.Println(string(buf[:len]))
			if string(buf[:len]) == "fail" {
				fmt.Println("the file doesn't exist")
				serverrec.Close()
				break
			}
			fmt.Println("receive filename:", filename)
			_, err = io.Copy(file, conn2)
			if err != nil {
				fmt.Println(err)
			}
			file.Close()
			serverrec.Close()
			fmt.Println("end down case")
			break
		}
		if triminput == "Q" {
			fmt.Println("quit")
			break
		}
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

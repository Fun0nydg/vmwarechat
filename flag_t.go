package main

import (
	"flag"
	"fmt"
)

func main() {
	ip := flag.String("ip", "", "")
	port := flag.Int("p", 80, "")
	fmt.Println("start")
	flag.Parse()
	fmt.Println(*ip, *port)
}

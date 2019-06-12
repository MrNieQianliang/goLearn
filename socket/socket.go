package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr,"Useage:%s ip-addr:",os.Args[0])
		os.Exit(1)
	}

	name := os.Args[0]
	addr := net.ParseIP(name)
	fmt.Println(addr)
}

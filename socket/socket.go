package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println(len(os.Args))

	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr,"%s",os.Args[0])
		os.Exit(1)
	}

	name := os.Args[0]
	addr := net.ParseIP(name)
	fmt.Println(addr)
}

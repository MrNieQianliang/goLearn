package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	go sayHello()
	go sayWorld()
	time.Sleep(100*time.Microsecond)
}

func sayHello()  {
	for i:=0;i<10 ;i++  {
		fmt.Println("Hello")
	}
}

func sayWorld()  {
	for i :=0 ;i<10 ;i++  {
		fmt.Println("Wrold")
	}

}

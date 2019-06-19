package main

import "fmt"

func main() {
	funcation1()
}

func funcation1()  {
	fmt.Println("start funcation 1")
	defer funcation2()
	//funcation2()
	fmt.Println("end funcation 1")
}

func funcation2()  {
	fmt.Println("start funcation 2")
	fmt.Println("end funcation 2")
}
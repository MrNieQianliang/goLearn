package main

import "fmt"

func main() {
	var text string = "nql"
	fmt.Println(&text)
	fmt.Println(*&text)
}

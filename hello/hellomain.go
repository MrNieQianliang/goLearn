package main

import "fmt"

func main() {

	fplus := func(x,y int) int { return x+y }

	fmt.Println(fplus(3,4))
}

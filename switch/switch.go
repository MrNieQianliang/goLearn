package main

import "fmt"

func main() {
	types := [] int{1,2,3,4,5,6,7}

	for x:=range types {
		switch x {
		case 1:
			fmt.Println("星期一")
		case 2:
			fmt.Println("星期二")
		case 3:
			fmt.Println("星期三")
		case 4:
			fmt.Println("星期四")
			fallthrough
		case 5:
			fmt.Println("星期五")
		case 6:
			fmt.Println("星期六")
		case 7:
			fmt.Println("星期日")
		}
	}
}

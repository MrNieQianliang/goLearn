package main

import "fmt"

func main() {
	//初始化切片
	var slice []int = make([]int,10)
	s := []int {1,2,3,4,5,6,7,8,99,0}
	for a:=0;a< len(slice);a++  {
		slice[a] = a;
	}
	//追加动态数组的元素
	slice = append(slice, 1,3,4,5,4)
	PrintSlice(slice)
	fmt.Println("输出另一个数组")
	PrintSlice(s)
}

func PrintSlice(s []int)  {
	for ss:=range s {
		fmt.Println(ss)
	}
}

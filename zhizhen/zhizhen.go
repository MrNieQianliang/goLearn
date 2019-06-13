package main

import "fmt"

func main() {
	var text string = "nql"
	point := &text
	*point = "nieqianliang"
	fmt.Println(&text)
	fmt.Println(*&text)

	var texts = [5][5]int {{11,12,13,14,15},{21,22,23,24,25},{31,32,33,34,35},{41,42,43,44,45},{51,52,53,54,55}}

	point_int := &texts[0][0]

	for i :=0; i < 5*5 ;i++  {
		fmt.Println(*(point_int)+i)
	}
}

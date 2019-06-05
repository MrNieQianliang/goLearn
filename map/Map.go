package main

import "fmt"

func main() {
	var maps map[string]string
	maps = make(map[string]string)
	maps ["1"] = "n"
	maps ["2"] = "q"
	maps ["3"] = "l"

	for m := range maps{
		fmt.Println(m)
	}
}

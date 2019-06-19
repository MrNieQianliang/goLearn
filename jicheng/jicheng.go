package main

import "fmt"

type Circle struct {
	radios float64
} 
func main() {
	circle := new(Circle)
	circle.radios = 100.0
	fmt.Println(circle.getArea())
}

func (c Circle) getArea() float64  {
	return 3.14*c.radios*c.radios;
}
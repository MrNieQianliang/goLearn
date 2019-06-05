package main

import "fmt"

type Phone interface {
	call()
	say()
}

type NokiaPhone struct {

}

type IPhone struct {
	
}

func (nokiaPhone NokiaPhone) call()  {
	fmt.Println("I am NokiaPhone")
}

func (iphone IPhone) call()  {
	fmt.Println("I am Iphone")
}

func (nokiaPhone NokiaPhone) say()  {
	fmt.Println("Hello NokiaPhone")
}

func (iphone IPhone) say()  {
	fmt.Println("Hello Iphone")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

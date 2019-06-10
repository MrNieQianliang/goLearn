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
	phone.say()

	fmt.Println(&phone)

	phone = new(IPhone)
	phone.call()
	phone.say()
	fmt.Println(&phone)
}

//Go语言使用接口的方式
//接口相当于规定公共方法，实现方法需要实现每一个抽象的方法
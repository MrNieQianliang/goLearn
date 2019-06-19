package main

import "fmt"

type Book struct {
	id int
	name string
	qrcode string
	title string
}
func main() {

	//使用结构体.成员名的方式访问
	var book1 Book
	book1.title = "go"
	book1.name = "go yuyan"
	book1.qrcode = "1234567"
	book1.id = 1

	fmt.Println(book1)

	fmt.Println(Book{1,"go1","12345","go yuyan"})

	fmt.Print(Book{id:90,qrcode:"321654",title:"go 语言"})
}

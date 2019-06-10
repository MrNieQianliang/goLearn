package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql" //导入第三方包
)

//定义接受数据的结构体
type User struct {
	name string
	age int
	email string
}


func main() {
	//打开连接数据库
	db,e := sql.Open("mysql","root:@tcp(localhost:3306)/MyTestData?charset=utf8")
	if (e!=nil) {
		fmt.Print("连接失败")
		fmt.Println(e)
		return
	}
	fmt.Println("连接数据库成功")

	_,e2 :=db.Query("select 1")
	if e2!=nil {
		fmt.Println(e2)
		return
	}

	rows,err := db.Query("select name,age,email from user")

	if err!=nil {
		fmt.Println(err)
		return
	}

	//循环读取数据
	for rows.Next(){
		user := new(User)
		row_text := rows.Scan(&user.name,&user.age,&user.email)
		if row_text!=nil {
			return
		}
		fmt.Printf("%s,%d,%s \n",user.name,user.age,user.email)
	}

	//插入数据库
	result,err1 := db.Exec("insert user(name,age,email)value (?,?,?)","nql","30","nql@qq.com")
	if err1!=nil {
		fmt.Println(err1)
		return
	}
	changerow,err2 := result.RowsAffected()
	if err2!=nil {
		fmt.Println(err2)
		return
	}

	fmt.Printf("%d 受影响数据",changerow)


}

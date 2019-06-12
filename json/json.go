package main

import (
	"encoding/json"
	"fmt"
)

//go语言处理json,json的Key必须是首字母大写，作为一个可导出的字段

type Server struct {
	Servername string
	Serverip string
}
type serverlice struct {
	Servers []Server
}


func main() {
	var s serverlice
	//json的字符串必须使用刀秋字符括起来才可以正常使用
	str := `{"servers":[{"servername":"zhengzhou_VPN","serverip":"127.0.0.1"},{"servername":"hangzhouzhou_VPN","serverip":"127.0.0.1"}]}`
	json.Unmarshal([]byte(str),&s)
	fmt.Println(s)

	var in interface{}
	json.Unmarshal([]byte(str),&in)
	fmt.Println(in)
	
}

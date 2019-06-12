package main

import (
	"fmt"
	"log"
	"net/http"
)

//一个简单的网络服务

func HelloServer(w http.ResponseWriter,req *http.Request)  {
	fmt.Println("Insert Hello charter")
	fmt.Fprintf(w,"hello,"+req.URL.Path[1:])
}

func WorldServer(w http.ResponseWriter,req *http.Request)  {
	fmt.Fprint(w,"Hello world")
}
func main() {
	http.HandleFunc("/hello",HelloServer)
	http.HandleFunc("/world",WorldServer)
	err:=http.ListenAndServe("localhost:8080",nil)
	if err !=nil {
		log.Fatal(err)
	}
}



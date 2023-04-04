package main

import (
	"fmt"
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("有一个http连接访问/say_hello")
	w.Write([]byte("hello"))
}

func goServer() {
	http.HandleFunc("/say_hello", SayHello)
	err := http.ListenAndServe(":8080", nil) //开启一个http服务
	if err != nil {
		log.Print("ListenAndServe: ", err)
		return
	}
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GinServerSayHello(c *gin.Context) {
	fmt.Println("有一个http连接访问/say_hello")
	c.String(http.StatusOK, "hello") //以字符串"hello"作为返回包
}

func ginServer() {
	engine := gin.Default() //生成一个默认的gin引擎
	engine.Handle(http.MethodGet, "/say_hello", GinServerSayHello)
	err := engine.Run(":8080") //使用8080端口号，开启一个web服务
	if err != nil {
		log.Print("server err: ", err.Error())
		return
	}
}

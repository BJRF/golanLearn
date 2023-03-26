package main

import "fmt"

// 全局变量
// 这里的全局变量如果要在包外访问，首字母需要大写，golang是以首字母大小写来区分对包外是否可见。
var globalStr string
var globalInt int

func initVarTest() {
	// 局部变量
	var localStr string
	var localInt int
	localStr = "local str"
	localInt = 123
	globalStr = "global str"
	globalInt = 123
	fmt.Printf("%s %d\n", localStr, localInt)
	fmt.Printf("%s %d\n", globalStr, globalInt)
}

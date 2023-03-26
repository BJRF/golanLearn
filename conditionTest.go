package main

import "fmt"

func conditionTest() {
	// := 这种方式可以直接初始化变量
	localStr := "case3"

	//基本if-else语句
	if localStr == "case3" {
		fmt.Printf("into case3 logic\n")
	} else {
		fmt.Printf("into no case3 logic\n")
	}

	var dic = map[string]int{
		"hjf":  100,
		"BJRF": 100,
	}
	//这里语句逻辑等同于下面的 if num, ok := dic["bjr"]; ok
	num, ok := dic["bjr"]
	fmt.Println(ok)
	if ok {
		fmt.Printf("bjr num %d", num)
	}
	//将表达式dic["bjr"]赋值给num,并判断语句是否抛出异常以bool值返回给ok
	if num, ok := dic["bjr"]; ok {
		fmt.Printf("bjr num %d\n", num)
	}
	if num, ok := dic["hjf"]; ok {
		fmt.Printf("hjf num %d\n", num)
	}

	//switch用法
	switch localStr {
	case "case1":
		fmt.Println("case1")
	case "case2":
		fmt.Println("case2")
	case "case3":
		fmt.Println("case3")
	default:
		fmt.Println("default")
	}
}

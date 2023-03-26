package main

import "fmt"

func testArraySliceDict() {
	// 初始化数组
	var strArray = [10]string{"a", "bc", "def"}

	/*
		切片的本质：
		1.它不是一个数组的指针，是一种数据结构体，用来操作数组内部元素
		2.在runtine/slice.go ,可以看出切片是种特殊的结构体
			type slice struct{
				*p
				len
				cap
			}
	*/
	/* 切片的创建：
	1.自动推导    slice:=[]int{1,2,3,4,5}
	2.slice:=make([]int,长度,容量)
	3.slice:=make([]int,长度) //创建切片时,没有指定容量的,容量=长度
	*/
	var sliceArray = make([]string, 0)
	// 初始化切片
	// 切片名称[low:high:max]
	// low:起始下标位置
	// high:结束下标位置 len=high-low
	// 容量cap: cap=max-low
	sliceArray = strArray[0:2]

	// 初始化字典并"构造函数"
	var dict1 = map[string]int{
		"hjf":  100,
		"BJRF": 100,
	}
	// 初始化字典不"构造函数"(感觉make像C++的new)
	var dict2 = make(map[string]int)
	dict2["AC"] = 100
	dict2["OC"] = 100

	//%v 只输出所有的值
	//%+v 先输出字段名字，再输出该字段的值
	//%#v 先输出结构体名字值，再输出结构体（字段名字+字段的值）
	fmt.Printf("%v\n", strArray)
	fmt.Printf("%v\n", sliceArray)
	fmt.Printf("%v\n", dict1["hjf"])
	fmt.Printf("%v\n", dict2)
}

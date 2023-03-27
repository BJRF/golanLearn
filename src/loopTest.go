package main

import "fmt"

func loopTest() {
	// 普通for循环
	for i := 0; i < 5; i++ {
		fmt.Printf("general for loop current i %d\n", i)
	}

	// 数组for循环
	var strArray = []string{"str0", "str1", "str2", "str3", "str4"}
	for i, str := range strArray {
		fmt.Printf("Array for loop i %d, str %s\n", i, str)
	}

	//切片for循环
	var sliceArray = make([]string, 0)
	sliceArray = strArray[0:3] //下标左开右闭
	for i, str := range sliceArray {
		fmt.Printf("slice for loop i %d, str %s\n", i, str)
	}

	//字典for循环
	var dic = map[string]int{
		"hjf":  100,
		"BRJF": 100,
	}
	for k, v := range dic {
		fmt.Printf("dict for loop key %s, value %d\n", k, v)
	}
}

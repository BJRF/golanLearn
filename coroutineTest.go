package main

import (
	"fmt"
	"time"
)

func a() {
	time.Sleep(3 * time.Second)
	fmt.Println("func a")
}

func b() {
	time.Sleep(2 * time.Second)
	fmt.Println("func b")
}

func c() {
	time.Sleep(1 * time.Second)
	fmt.Println("func c")
}

func coroutineTest() {
	a()
	b()
	c()
	go a()
	go b()
	go c()
	time.Sleep(5 * time.Second) //最后不加sleep则不会输出go a b c
}

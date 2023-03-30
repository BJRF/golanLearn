package main

import (
	"fmt"
	"sync"
)

/*
panic抓取error实践
*/
func panicTest() {
	fmt.Println("c")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("e")
	}()
	f()              //开始调用f
	fmt.Println("f") //这里开始下面代码不会再执行
}

func f() {
	fmt.Println("a")
	panic("异常信息")
	fmt.Println("b") //这里开始下面代码不会再执行
	fmt.Println("f")
}

func errorTest() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("e")
	}()
	s := make([]string, 3)
	fmt.Println(s[5])
}

// 抓取协程中的panic实践
func withGoroutine(opts ...func() error) (err error) {
	var wg sync.WaitGroup
	for _, opt := range opts {
		wg.Add(1)
		// 开启goroutine，做并行处理
		go func(handler func() error) {
			defer func() {
				// 协程内部捕获panic
				if e := recover(); e != nil {
					fmt.Printf("recover:%v\n", e)
				}
				wg.Done()
			}()
			e := handler() // 真正的逻辑调用
			// 因为handler执行了panic，所以下面阻塞了，并不会调用。
			fmt.Printf("返回报错:%v\n", e)
			// 取第一个报错的handler调用的错误返回
			if err == nil && e != nil {
				err = e
				fmt.Printf("返回第一个报错:%v\n", err)
			}
		}(opt) // 将goroutine的函数逻辑通过封装成的函数变量传入
	}

	wg.Wait() // 等待所有的协程执行完

	return
}

func recoverTest() {
	handler1 := func() error {
		panic("handler1 fail ")
		return fmt.Errorf("return handler1 err")
	}

	handler2 := func() error {
		panic("handler2 fail")
		return fmt.Errorf("return handler2 err")
	}

	err := withGoroutine(handler1, handler2) // 并发执行handler1和handler2两个任务，返回第一个报错的任务错误
	if err != nil {
		fmt.Printf("err is:%v", err)
	}
}

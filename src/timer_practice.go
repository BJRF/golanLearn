package main

import (
	"fmt"
	"time"
)

// Timer数据结构
//
//	type Timer struct {
//		C <-chan Time
//		r runtimeTimer
//	}
func TimerTest() {
	//创建Timer
	timer := time.NewTimer(10 * time.Second) //设置超时时间2s

	//停止Timer
	stopRes := timer.Stop()
	fmt.Printf("StopRes:%v\n", stopRes)

	//已停止的TImer可以重置Timer
	resetRes := timer.Reset(time.Second * 3)
	<-timer.C
	fmt.Printf("reset Timer%v\n", resetRes)

	//time.AfterFunc参数为超时时间d和一个具体的函数f，返回一个Timer的指针，
	//作用在创建出timer之后，在当前goroutine，等待一段时间d之后，将执行f。
	duration := time.Duration(1) * time.Second
	f := func() {
		fmt.Println("f has been called after 1s by time.AfterFunc")
	}
	_ = time.AfterFunc(duration, f)
	time.Sleep(2 * time.Second)

	ch := make(chan string)

	//after函数经过时间短d之后会返回timer里的管道，并且这个管道会在经过时段d之后写入数据，
	//调用这个函数，就相当于实现了定时器。 一般time.After会配合select一起使用
	go func() {
		time.Sleep(time.Second * 2)
		ch <- "test"
	}()
	select {
	case val := <-ch:
		fmt.Printf("val is %s\n", val)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout!!!")
	}
}

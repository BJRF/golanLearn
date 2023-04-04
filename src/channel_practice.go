package main

import (
	"fmt"
	"time"
)

/*
总结：
是否有缓冲对生产者生效，若有缓冲空间则不阻塞，无缓冲空间则阻塞
对消费者来说如果没有消息则一直阻塞
*/

var chSync chan int
var chAsyn chan int

// 无缓冲测试
func syncChannelTestA() {
	time.Sleep(3 * time.Second)
	a := 5
	//由于是无缓冲的，A的消息未被消费者拿到，所以阻塞
	chSync <- a
	fmt.Println("out of channelTestA")
}
func syncChannelTestB() {
	time.Sleep(1 * time.Second)
}

// 有缓冲测试
func AsynChannelTestA() {
	time.Sleep(1 * time.Second)
	//time.Sleep(3 * time.Second)
	a := 5
	//由于是有缓冲通道，A在消息直接进入缓冲，不阻塞
	chAsyn <- a
	fmt.Println("out of channelTestA")
}
func AsynChannelTestB() {
	time.Sleep(3 * time.Second)
	//time.Sleep(1 * time.Second)
	//这里B会阻塞，直到通道中有消息可消费
	fromA := <-chAsyn
	fmt.Println("channelTestA is ", fromA)
}

func channelTest() {
	//全局变量定义后在局部初始化分配内存
	//无缓冲
	chSync = make(chan int)
	go syncChannelTestA()
	go syncChannelTestB()
	time.Sleep(5 * time.Second)

	//有缓冲
	chAsyn = make(chan int, 1)
	go AsynChannelTestA()
	go AsynChannelTestB()
	time.Sleep(5 * time.Second)

	fmt.Println("out of main")
}

package main

import (
	"fmt"
	"time"
)

func SelectTest() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		for i := 0; i < 3; i++ {
			select {
			case v := <-ch1:
				fmt.Printf("Received from ch1, val = %d\n", v)
			case v := <-ch2:
				fmt.Printf("Received from ch2, val = %d\n", v)
			default:
				fmt.Println("default!!!")
			}
			time.Sleep(time.Second)
		}
	}()
	ch1 <- 1
	time.Sleep(time.Second)
	ch2 <- 2
	time.Sleep(4 * time.Second)
	// 结束后因为select没有数据，调用default
}

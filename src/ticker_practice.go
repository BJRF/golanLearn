package main

import (
	"fmt"
	"time"
)

func WatchTickerTest() chan struct{} {
	ticker := time.NewTicker(1 * time.Second)

	ch := make(chan struct{})
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("watch!!!")
			case <-ch:
				fmt.Println("Ticker Stop!!!")
				return
			}
		}
	}(ticker)
	return ch
}

// TickerTest
// Ticker对象的字段和Timer是一样的，也包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间，
// 根据这个管道消息来触发事件，但是ticker只要定义完成，就从当前时间开始计时，每隔固定时间都会触发，
// 只有关闭Ticker对象才不会继续发送时间消息
func TickerTest() {
	ch := WatchTickerTest()
	time.Sleep(5 * time.Second)
	ch <- struct{}{}
	close(ch)
}

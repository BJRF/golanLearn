package main

import (
	"context"
	"fmt"
	"time"
)

func ContextTest() {

	//WithCancel用法
	//context.WithCancel 函数是一个取消控制函数，只需要一个context作为参数，
	//能够从 context.Context 中衍生出一个新的子context和取消函数CancelFunc，
	//通过将这个子context传递到新的goroutine中来控制这些goroutine的关闭，
	//一旦我们执行返回的取消函数CancelFunc，当前上下文以及它的子上下文都会被取消，所有的 Goroutine 都会同步收到取消信号。
	//ctx, cancel := context.WithCancel(context.Background())

	//WithDeadline用法
	//context.WithDeadline也是一个取消控制函数，方法有两个参数，第一个参数是一个context，第二个参数是截止时间，
	//同样会返回一个子context和一个取消函数CancelFunc。在使用的时候，没有到截止时间，我们可以通过手动调用CancelFunc来取消子context，
	//控制子goroutine的退出，如果到了截止时间，我们都没有调用CancelFunc，子context的Done()管道也会收到一个取消信号，用来控制子goroutine退出。
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(4*time.Second)) // 设置超时时间4当前时间4s之后

	//WithTimeout用法
	//context.WithTimeout和context.WithDeadline的作用类似，都是用于超时取消子context，
	//只是传递的第二个参数有所不同，context.WithTimeout传递的第二个参数不是具体时间，而是时间长度。
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)

	//WithValue用法
	//context.WithValue 函数从父context中创建一个子context用于传值，函数参数是父context，key，val键值对。返回一个context。
	//项目中这个方法一般用于上下文信息的传递，比如请求唯一id，以及trace_id等，用于链路追踪以及配置透传。
	//ctx := context.WithValue(context.Background(), "name", "zhangsan")
	//go func1(ctx)
	//func func1(ctx context.Context) {
	//	fmt.Printf("name is: %s", ctx.Value("name").(string))
	//}

	go Watch(ctx, "goroutine1")
	go Watch(ctx, "goroutine2")

	time.Sleep(6 * time.Second) // 让goroutine1和goroutine2执行6s
	fmt.Println("end watching!!!")
	cancel() // 通知goroutine1和goroutine2关闭
	time.Sleep(1 * time.Second)
}

func Watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s exit!\n", name) // 主goroutine调用cancel后，会发送一个信号到ctx.Done()这个channel，这里就会收到信息
			return
		default:
			fmt.Printf("%s watching...\n", name)
			time.Sleep(time.Second)
		}
	}
}

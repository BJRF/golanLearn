package main

import (
	"fmt"
	"time"
)

type Task struct {
	f func() error // 具体的任务逻辑
}

func NewTask(funcArg func() error) *Task {
	return &Task{
		f: funcArg,
	}
}

type Pool struct {
	WorkerNum int        // goroutine数量
	JobCh     chan *Task // 用于worker取任务
}

func NewPool(workerNum int, taskNum int) *Pool {
	return &Pool{
		// worker数
		WorkerNum: workerNum,
		// 设置了channel的大小，也就是限制了task的最大数量
		JobCh: make(chan *Task, taskNum),
	}
}

func (p *Pool) Size() int {
	return p.WorkerNum
}

// worker 这个角色就是协程池里干活的goroutine
func (p *Pool) worker(i int) {
	// worker只做一件事，从管道里取出task，并且执行
	for task := range p.JobCh {
		if err := task.f(); err != nil {
			fmt.Printf("worker %d handle fail: %v\n", i+1, err)
			return
		}
		fmt.Printf("worker %d handle success\n", i+1)
	}
}

// AddTask 往协程池里面添加任务
func (p *Pool) AddTask(task *Task) {
	p.JobCh <- task
}

// Run 协程池跑起来，需要指定数量的worker干活
func (p *Pool) Run() {
	// 这里只创建指定数量的worker来干活
	// 这里限制的worker的数量
	for i := 0; i < p.Size(); i++ {
		go p.worker(i)
	}
}

func GoroutinePoolTest() {
	p := NewPool(3, 10)
	p.Run()
	task := NewTask(func() error {
		fmt.Printf("I am Task\n")
		return nil
	})
	for i := 0; i < 100; i++ {
		p.AddTask(task)
		fmt.Printf("task num:%v\n", i+1)
	}
	time.Sleep(3 * time.Second)
}

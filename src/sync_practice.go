package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

/*
channel锁
*/

// ch <- true和<- ch就相当于一个锁，将 *num = *num + 1这个操作锁住了。
// 因为ch管道的容量是1，在每个add函数里都会往channel放置一个true，
// 直到执行完+1操作之后才将channel里的true取出。由于channel的size是1，
// 所以当一个goroutine在执行add函数的时候，其他goroutine执行add函数，
// 执行到ch <- true的时候就会阻塞，*num = *num + 1不会成功，直到前一个
// +1操作完成，<-ch，读出了管道的元素，这样就实现了并发安全
func add(ch chan bool, num *int) {
	ch <- true
	*num = *num + 1
	<-ch
}

func ChannelSyncLockPractice() {
	// 创建一个size为1的channel
	ch := make(chan bool, 1)

	var num int
	for i := 0; i < 100; i++ {
		go add(ch, &num)
	}

	time.Sleep(2)
	fmt.Println("num 的值：", num)
}

/*
sync.WaitGroup()实践
*/
//可以用sync包下的WaitGroup来实现，Go语言中可以使用sync.WaitGroup来实现并发任务的同步以及协程任务等待。
//sync.WaitGroup是一个对象，里面维护者一个计数器，并且通过三个方法来配合使用
//- (wg * WaitGroup) Add(delta int)       计数器加delta
//- (wg *WaitGroup) Done()                    计数器减1
//- (wg *WaitGroup) Wait()                     会阻塞代码的运行，直至计数器减为0。

var wg sync.WaitGroup

func waitGroupGoroutine() {
	defer wg.Done()
	fmt.Println("waitGroupGoroutine!")
}
func waitGroupTest() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go waitGroupGoroutine()
	}
	wg.Wait()
	fmt.Println("end!!!")
}

/*
sync.Once()实践
*/
// 在我们写项目的时候，程序中有很多的逻辑只需要执行一次，最典型的就是项目工程里配置文件的加载，我们只需要加载一次即可，
// 让配置保存在内存中，下次使用的时候直接使用内存中的配置数据即可。这里就要用到sync.Once。
// sync.Once，并且线程安全。sync.Once最大的作用就是延迟初始化，
// 对于一个使用sync.Once，而是在第一次用的它的时候才会初始化，并且只初始化这一次，
// 初始化之后驻留在内存里，这就非常适合我们之前提到的配置文件加载场景，设想一下，如果是在程序刚开始就加载配置，
// 若迟迟未被使用，则既浪费了内存，又延长了程序加载时间，而sync.Once解决了这个问题。
// 声明配置结构体Config
// 类似于单例模式
type syncOnceConfig struct{}

var instance syncOnceConfig
var once sync.Once // 声明一个sync.Once变量

func getOnceConfig() syncOnceConfig {
	var syncOnceConfig syncOnceConfig
	return syncOnceConfig
}

// 获取配置结构体
func InitConfig() *syncOnceConfig {
	once.Do(func() {
		fmt.Println("初始化成功")
		instance = getOnceConfig()
	})
	return &instance
}

/*
互斥锁
*/

//一个互斥锁只能同时被一个 goroutine 锁定，其它 goroutine 将阻塞直到互斥锁被解锁才能加锁成功
//sync.Mutex在使用的时候要注意：对一个未锁定的互斥锁解锁将会产生运行时错误

var mutexTestNum int

func AddMutex(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock() // 加锁
	defer func() {
		wg.Done()   // 计数器减1
		mu.Unlock() // 解锁
	}()
	mutexTestNum += 1
}

func MutexTest() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(10) // 开启10个goroutine，计数器加10
	for i := 0; i < 10; i++ {
		go AddMutex(&wg, &mu)
	}
	wg.Wait() // 等待所有协程执行完
	fmt.Println(mutexTestNum)
}

/*
读写锁
*/

//读写锁就是将读操作和写操作分开，可以分别对读和写加锁，一般用在大量读操作、少量写操作的情况
//1. 同时只能有一个 goroutine 能够获得写锁定。
//2. 同时可以有任意多个 goroutine 获得读锁定。
//3. 同时只能存在写锁定或读锁定（读和写互斥）。
//可以多个goroutine同时读，但是只有一个goroutine能写，
//共享资源要么在被一个或多个goroutine读取，要么在被一个goroutine写入， 读写不能同时进行

var rWMutexNum = 0

func RWMutexTest() {
	var mr sync.RWMutex
	for i := 1; i <= 3; i++ {
		go RWMutexTestwrite(&mr, i)
	}
	for i := 1; i <= 3; i++ {
		go RWMutexTestread(&mr, i)
	}
	time.Sleep(time.Second)
	fmt.Println("final count:", rWMutexNum)
}

func RWMutexTestread(mr *sync.RWMutex, i int) {
	fmt.Printf("goroutine%d reader start\n", i)
	mr.RLock()
	fmt.Printf("goroutine%d reading count:%d\n", i, rWMutexNum)
	time.Sleep(time.Millisecond)
	mr.RUnlock()
	fmt.Printf("goroutine%d reader over\n", i)
}

func RWMutexTestwrite(mr *sync.RWMutex, i int) {
	fmt.Printf("goroutine%d writer start\n", i)
	mr.Lock()
	rWMutexNum++
	fmt.Printf("goroutine%d writing count:%d\n", i, rWMutexNum)
	time.Sleep(time.Millisecond)
	mr.Unlock()
	fmt.Printf("goroutine%d writer over\n", i)
}

/*
死锁场景1
*/

// DeadlockScenario1 Lock/Unlock不成对
// 这里mu sync.Mutex当作参数传入到函数copyMutex，锁进行了拷贝，不是原来的锁变量了，
// 那么一把新的锁，在执行mu.Lock()的时候应该没问题。这就是要注意的地方，
// 如果将带有锁结构的变量赋值给其他变量，锁的状态会复制。所以多锁复制后的新的锁拥有了原来的锁状态，
// 那么在copyMutex函数内执行mu.Lock()的时候会一直阻塞，因为外层的main函数已经Lock()了一次，
// 但是并没有机会Unlock()，导致内层函数会一直等待Lock()，而外层函数一直等待Unlock()，这样就造成了死锁
// 所以在使用锁的时候，我们应当尽量避免锁拷贝，并且保证Lock()和Unlock()成对出现，
// 没有成对出现容易会出现死锁的情况，或者是Unlock 一个未加锁的Mutex而导致 panic。
func DeadlockScenario1() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	copyMutex(mu)
}

func copyMutex(mu sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("ok")
}

/*
死锁场景2
*/

// DeadlockScenario2 互相等待
func DeadlockScenario2() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		mu1.Lock()
		defer mu1.Unlock()
		time.Sleep(1 * time.Second)

		mu2.Lock()
		defer mu2.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu2.Lock()
		defer mu2.Unlock()
		time.Sleep(1 * time.Second)
		mu1.Lock()
		defer mu1.Unlock()
	}()
	wg.Wait()
}

/*
解决go的map线程不安全的问题
*/

// 1、使用读写锁

var MySyncMapTestMap = make(map[string]int)
var mySyncMapMutex sync.Mutex

func getVal(key string) int {
	return MySyncMapTestMap[key]
}

func setVal(key string, value int) {
	MySyncMapTestMap[key] = value
}

func MySyncMapTest() {
	wg := sync.WaitGroup{}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer func() {
				wg.Done()
				mySyncMapMutex.Unlock()
			}()
			key := strconv.Itoa(num)
			mySyncMapMutex.Lock()
			setVal(key, num)
			fmt.Printf("key=:%v,val:=%v\n", key, getVal(key))
		}(i)
	}
	wg.Wait()
}

// 2、go1.19后可以使用sync内自带的安全map

func SyncMapTest() {
	var m sync.Map
	// 1. 写入
	m.Store("name", "hjf")
	m.Store("age", 18)

	// 2. 读取
	age, _ := m.Load("age")
	fmt.Println(age.(int))

	// 3. 遍历
	// Range函数的参数是一个有两个interface{}类型参数，返回值为bool类型的的函数，这里的意思是函数的参数是一个函数
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key is:%v, val is:%v\n", key, value)
		return true
	})

	// 4. 删除
	m.Delete("age")
	age, ok := m.Load("age")
	fmt.Println(age, ok)

	// 5. 读取或写入
	m.LoadOrStore("name", "hjf")
	name, _ := m.Load("name")
	fmt.Println(name)
}

/*
sync.Atomic
*/
//使用方式：通常mutex用于保护一段执行逻辑，而atomic主要是对变量进行操作
//底层实现：mutex由操作系统调度器实现，而atomic操作有底层硬件指令支持，保证在cpu上执行不中断。所以atomic的性能也能随cpu的个数增加线性提升

func AtomicTest() {
	var sum int32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&sum, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("sum is %d\n", sum)
}

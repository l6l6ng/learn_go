package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//这个示例程序展示如何使用atomic包里的
//Store和Load类函数来提供对数值类型
//的安全访问

var (
	//shutdown是通知正在执行的goroutine停止工作的标志
	shutdown int64

	//wg 用来等待程序结束
	wg sync.WaitGroup
)

// main是所有Go程序的入口
func main() {
	//计数加2，表示要等待两个goroutine
	wg.Add(2)

	//创建两个goroutine
	go doWork("A")
	go doWork("B")

	//给定goroutine执行的时间
	time.Sleep(1 * time.Second)

	//该停止工作了，安全地设置shutdown标志
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	//等待goroutine结束
	wg.Wait()
}

//doWork用来模拟执行工作的goroutine,
//检测之前的shutdown 标志来决定是否提前终止
func doWork(name string) {
	//在函数退出时调用Done来通知main函数工作已经完成
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Word\n", name)
		time.Sleep(250 * time.Millisecond)

		//要停止工作了吗？
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}

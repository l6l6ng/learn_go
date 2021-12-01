// 这个示例程序展示如何使用有缓冲的通道和固定数目的goroutine来处理一堆工作
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numberGoroutine = 4 //要使用的goroutine的数量
const taskLoad = 10       //要处理的工作的数量

// wg 用来等待程序完成
var wg sync.WaitGroup

func init() {
	//初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {
	//创建一个有缓冲的通道来处理工作
	tasks := make(chan string, taskLoad)

	//启动goroutine来处理工作
	wg.Add(numberGoroutine)
	for gr := 1; gr <= numberGoroutine; gr++ {
		go worker(tasks, gr)
	}

	//增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	close(tasks)

	wg.Wait()
}

// worker 作为goroutine启动来处理从有缓存的通道传入的工作
func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		//等待分配工作
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Workder: %d : Shutting Down\n", worker)
			return
		}

		//显示我们开始工作了
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		//随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		//显示我们完成了工作
		fmt.Printf("Worker: %d ： Completed %s \n", worker, task)
	}
}

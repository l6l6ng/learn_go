package main

import (
	"fmt"
)

type Queue struct {
	maxSize int
	array [5]int
	front int
	rear int
}

//添加数据到队列
func (this *Queue) AddQueue(val int) (error) {
	//先判断队列是否已满
	if this.rear == this.maxSize - 1 {
		return errors.New("queue full")
	}

	this.rear++//rear 后移
	this.array[this.rear] = val
	return 
}

//显示队列
func (this *Queue) ShowQueue() {
	//this.front 不包含队首的元素
	for i := this.front + 1;i<=this.rear;i++{
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
}

//编写一个主函数测试，

func main(){
	queue = Queue{
		maxSize: 5,
		
	}
}
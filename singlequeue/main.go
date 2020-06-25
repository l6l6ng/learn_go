package main

import (
	"fmt"
)

type Queue struct {
	maxSize 4
	array [4]int
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
	
}
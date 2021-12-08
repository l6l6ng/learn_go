// 广度优先搜索指出是否有从A到C的路径
// 如有广度优先搜索可找出最短路径
package main

import (
	"container/list"
	"fmt"
)

func main() {
	m := make(map[string][]string)
	m["a"] = []string{"d", "g", "b"}
	m["d"] = []string{"e", "e1"}
	m["b"] = []string{"h", "b1"}
	m["g"] = []string{"f", "i"}
	m["e"] = []string{}
	m["e1"] = []string{"c"}
	m["h"] = []string{}
	m["b1"] = []string{}
	m["f"] = []string{}
	m["i"] = []string{}
	m["c"] = []string{}

	search(m)
}

func search(m map[string][]string) {
	l := list.New()
	push(l, m["a"])

	searched := make([]interface{}, 0) //记录已经搜索过的元素

	// todo 判断队列是否取空
	for l.Len() > 0 {
		p := pop(l).(string) //取一个元素
		if !inArray(p, searched) {
			searched = append(searched, p)

			if check(p) {
				fmt.Printf("find %v\n", p)
				return
			}

			//未找到把下级关系放入队列
			push(l, m[p])
		}
	}
}

// push 入队
func push(l *list.List, s []string) {
	for _, v := range s {
		l.PushBack(v)
	}
}

// pop 出队
func pop(l *list.List) interface{} {
	l1 := l.Front()
	l.Remove(l1)
	return l1.Value
}

// inArray 排队元素是否在切片或数组中
func inArray(p interface{}, ary []interface{}) bool {
	for _, v := range ary {
		if v == p {
			return true
		}
	}
	return false
}

// 是否是c 表示找到了目标元素
func check(v interface{}) bool {
	if v == `c` {
		return true
	}
	return false
}

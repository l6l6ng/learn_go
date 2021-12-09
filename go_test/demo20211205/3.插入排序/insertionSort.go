package main

import "fmt"

func main() {
	var ary = []int{3, 7, 100, 99, 87, 45, 23, 87, 99, 23, 21, 0, 1, 46}
	fmt.Println(insertionSort(ary))
}

// insertionSort 插入排序
func insertionSort(ary []int) []int {
	len := len(ary)

	for i := 1; i < len; i++ {
		j := i
		for j > 0 && ary[j-1] > ary[j] {
			ary[j-1], ary[j] = ary[j], ary[j-1]
			j--
		}
	}
	return ary
}

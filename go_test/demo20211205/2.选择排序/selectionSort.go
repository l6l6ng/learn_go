package main

import "fmt"

var ary = []int{3, 7, 100, 99, 87, 45, 23, 87, 99, 23, 21, 0, 1, 46}

func main() {
	fmt.Println(selectionSort(ary))
}

func selectionSort(ary []int) []int {
	len := len(ary)
	for i := 0; i < len-1; i++ {
		min := i
		for j := i + 1; j < len; j++ {
			if ary[j] < ary[min] {
				min = j
			}
		}
		ary[i], ary[min] = ary[min], ary[i]
	}
	return ary
}

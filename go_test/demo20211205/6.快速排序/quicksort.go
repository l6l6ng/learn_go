package main

import "fmt"

var ary = []int{8, 99, 191, 3, 23, 45, 8, 17, 12, 45, 2, 4, 56, 76, 11, 100, 12}

func main() {
	fmt.Println(quicksort(ary))
}

func quicksort(ary []int) []int {
	if len(ary) < 2 {
		return ary
	}
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(ary); i++ {
		if ary[i] < ary[0] {
			left = append(left, ary[i])
		} else {
			right = append(right, ary[i])
		}
	}

	return append(append(quicksort(left), ary[0]), quicksort(right)...)
}

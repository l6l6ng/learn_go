package main

import "fmt"

var ary = []int{1, 2, 3, 4, 5, 6, 19, 200}

func main() {
	fmt.Println(binarySearch(ary, 2001))
}

func binarySearch(ary []int, item int) int {
	low := 0
	high := len(ary) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := ary[mid]
		if guess == item {
			return mid
		} else if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

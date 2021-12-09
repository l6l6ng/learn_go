package main

import "fmt"

func main() {
	var ary = []int{3, 7, 100, 99, 87, 45, 23, 87, 99, 23, 21, 0, 1, 46}
	fmt.Println(bubbleSort(ary))
}

// bubbleSort 冒泡排序
func bubbleSort(sli []int) []int {
	len := len(sli)
	for i := len - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}
	return sli
}

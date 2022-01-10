package main

import "fmt"

func main() {
	var a = []int{3, 34, 5, 65, 44, 1, 2, 4}

	fmt.Println(insertionSort2(a))
}

func insertionSort2(arr []int) []int {
	for i := range arr {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
	return arr
}

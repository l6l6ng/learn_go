package main

import "fmt"

func main() {
	a := []int{1, 1, 3, 2, 3, 4, 4, 5, 5}
	n := removeDuplicates(a)
	fmt.Println(n)
}

func removeDuplicates(nums []int) int {
	a := make(map[int]int)
	k := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := a[nums[i]]; !ok {
			a[nums[i]] = i
			nums[k] = nums[i]
			k++
		}
	}
	return len(a)
}

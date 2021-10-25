package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 1}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	len := len(nums)
	for i := 0; i < len; i++ {
		for k := i + 1; k < len; k++ {
			if nums[i] == nums[k] {
				return true
			}
		}
	}
	return false
}

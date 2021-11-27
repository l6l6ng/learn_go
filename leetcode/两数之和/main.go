package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	a := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n := target - nums[i]
		if _, ok := a[n]; !ok {
			a[nums[i]] = i
		} else {
			return []int{a[n], i}
		}
	}
	return nil
}

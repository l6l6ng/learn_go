package main

import "fmt"

func removeElement(nums []int, target int) int {
	n := 0
	for _, v := range nums {
		if v != target {
			n++
		}
	}
	return n
}

func removeElement2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != target {
			fmt.Println(nums)
			nums[i], nums[j] = nums[j], nums[i]
			fmt.Println(nums)
			j++
		}
	}
	return j
}

func main() {
	nums := []int{4, 6, 9, 2, 3, 1, 6, 8, 2, 9, 1, 1}
	fmt.Println(removeElement(nums, 9))
	fmt.Println(removeElement2(nums, 9))
}

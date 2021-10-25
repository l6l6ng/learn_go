package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	len := len(nums)
	k = k % len
	if len > 1 {
		for i := 0; i < len/2; i++ {
			nums[i], nums[len-1-i] = nums[len-1-i], nums[i]
		}

		for i := 0; i < k/2; i++ {
			nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
		}

		for i := k; i < (len+k)/2; i++ {
			nums[i], nums[len-1-i+k] = nums[len-1-i+k], nums[i]
		}
	}
}

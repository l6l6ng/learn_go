package main

import "fmt"

func main() {
	//nums := []int{1, 0, 1, 0, 12, 3}
	nums := []int{0, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	i := len(nums) - 1
	for m := 0; m <= i; m++ {
		for j := 0; j < i; j++ {
			if nums[j] == 0 {
				for k := j; k < i; k++ {
					nums[k], nums[k+1] = nums[k+1], nums[k]
				}
				i--
			}
		}
	}
}

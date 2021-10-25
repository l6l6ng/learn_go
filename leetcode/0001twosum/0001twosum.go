package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}

func twoSum2(numbers []int, target int) []int {
	// write code here
	len := len(numbers)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{4, 4, 3, 9, 2, 2, 7}, 8))
	fmt.Println(twoSum2([]int{1, 5, 3, 9, 2, 6, 7}, 7))
}

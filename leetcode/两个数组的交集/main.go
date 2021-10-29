package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3}
	fmt.Println(intersect(a, b))
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	list := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j += 1
		} else if nums1[i] < nums2[j] {
			i += 1
		} else {
			list = append(list, nums1[i])
			i += 1
			j += 1
		}
	}
	return list
}

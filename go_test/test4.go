package main

import (
	_ "crypto/md5"
	"fmt"
)

func main() {
	//str := []byte("123456")
	//fmt.Println(str)
	//md5Str := fmt.Sprintf("%x", md5.Sum(str))
	//fmt.Println(md5Str)

	//str2 := []byte(123456)
	//md5Str2 := fmt.Sprintf("%x", md5.Sum(str))
	//fmt.Println(md5Str2)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	nums = append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
}

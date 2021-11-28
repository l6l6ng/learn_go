package main

import "fmt"

func main() {
	slice1 := [][]int{{10}, {20, 30}}
	fmt.Printf("slice1:%p\n", slice1)
	slice2 := slice1[0]
	slice2 = append(slice2, 200)
	fmt.Printf("slice1:%p\n", slice1)

	fmt.Println(slice2)
}

//s2 修改切片的值
func s2(slice2 []int) {
	slice2[0] = 100
	//slice2 = append(slice2, 200)
}

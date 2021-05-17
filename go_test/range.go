package main 

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	for k,i := range arr {
		fmt.Println(k, i)
	}
}
package main

import "fmt"

var ary = [...]float32{1, 2, 3, 4, 5, 6}

func main() {
	a := ary[2]
	fmt.Println(a)

	for k, v := range ary {
		fmt.Println(k, v)
	}

	var ary2 = [...]int{1, 2, 3, 4}
	for k1, v2 := range ary2 {
		fmt.Println(k1, v2)
	}
}

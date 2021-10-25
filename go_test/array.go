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

	//for i := 7.0000; i >= 0; i-- {
	//	fmt.Println(i/7)
	//}
	//s := 0.1
	//for i := 1; i <= 1000; i++ {
	//	s += 0.1
	//}
	fmt.Println(7.0/8)
	fmt.Println(6.0/8)
	fmt.Println(4.0/8)
	fmt.Println(2.0/8)
}

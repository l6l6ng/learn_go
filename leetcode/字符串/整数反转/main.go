package main

import "fmt"

func main() {
	x := -12346
	fmt.Println(reverse(x))

}

func reverse(x int) int {
	a := 0
	for x != 0 {
		y := x % 10
		a += a*10 + y
		x = x % 10
	}
	return a
}

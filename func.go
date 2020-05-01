package main

import "fmt"

func f(a, b int) (int, int) {
	 a, b = b, a
	return a, b
}

func main(){
	for m, n := 1, 2; m < 10; m++ {
		fmt.Println(f(m,n))
	}
}
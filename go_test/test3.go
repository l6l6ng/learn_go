package main

import "fmt"
var a int

func main() {
	b := ff1()
	fmt.Println(b)
	fmt.Println(a)
}

func ff1() int {
	a = 1
	return a
}

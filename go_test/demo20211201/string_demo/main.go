package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "abc"
	s1 := str[0] - 'z'
	s2 := 'a' - 'z'
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println(reflect.TypeOf(s2))
}

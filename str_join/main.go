package main

import (
	"fmt"
	"strings"
)

func main() {
	str := strings.Join([]string{"a1","b2","c3"},"--")
	fmt.Println(str)
	ary := strings.Split(str,"--")
	fmt.Println(ary)
}

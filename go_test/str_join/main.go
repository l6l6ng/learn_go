package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	str := strings.Join([]string{"a1","b2","c3"},"--")
	fmt.Println(str)
	ary := strings.Split(str,"--")
	fmt.Println(ary)
	path := "/Users/liulong/person/code/learn_go/str_join/main.go"
	dir := filepath.Dir(path)
	fmt.Println(dir)
}

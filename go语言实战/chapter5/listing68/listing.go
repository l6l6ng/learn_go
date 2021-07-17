package main

import (
	"fmt"
	"learn_go/go语言实战/chapter5/counters"
	"reflect"
)

func main() {
	counter := counters.New(10)
	fmt.Printf("Counter: %d\n", counter)
	fmt.Println(reflect.TypeOf(counter))
}

package main

import "fmt"

func main() {
	var i []interface{}
	i = append(i, 1, "a", []string{"a", "b"})
	for _, value := range i {
		if v, ok := value.([]string); ok {
			fmt.Println(v)
		}
	}
}

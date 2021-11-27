package main

import "fmt"

func main() {
	s := []string{"h", "e", "l", "l", "o"}
	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []string) {
	len := len(s)
	for i := 0; i < len/2; i++ {
		s[i], s[len-1-i] = s[len-1-i], s[i]
	}
}

package main

import "fmt"

func hannota(n int, a, b, c string) {
	if n == 1 {
		fmt.Printf("将编号为%d的圆盘从%s移动到%s\n", n, a, c)
	} else {
		hannota(n-1, a, c, b)
		fmt.Printf("将编号为%d的圆盘从%s移动到%s\n", n, a, c)
		hannota(n-1, b, a, c)
	}
}

var a = "A"
var b = "B"
var c = "C"

func main() {
	hannota(5, a, b, c)
}

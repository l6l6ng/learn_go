//这个示例程序展示如何使用io.Reader和io.Writer接口
//写一个简单版本的curl程序
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	fmt.Println(os.Args)
	os.Exit(-1)
	if len(os.Args) != 2 {
		fmt.Println("Usage:./example2<url>")
		os.Exit(-1)
	}
}

func main() {
	//从web服务器得到相应
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	//从body复制到Stdout
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

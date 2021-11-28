package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var strUrl string = "https://juejin.cn/post/7035160857633882120"

func main() {
	for i := 0; i < 1; i++ {
		go func(i int) {
			resp, err := http.Get(strUrl)
			if err != nil {
				log.Println("error", i)
			}

			defer resp.Body.Close()

			res, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(resp.StatusCode)
			}
			fmt.Printf("%s\n", res)
		}(i)
	}
}

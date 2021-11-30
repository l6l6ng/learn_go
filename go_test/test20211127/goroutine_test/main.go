package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var strUrl string = "https://juejin.cn/post/7035160857633882120"

func main() {
	total := 10
	var wg = sync.WaitGroup{}
	wg.Add(total)

	for i := 0; i < total; i++ {
		go func(i int) {
			defer wg.Done()
			resp, err := http.Get(strUrl)
			if err != nil {
				log.Println("error", i)
			}

			defer resp.Body.Close()

			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(resp.StatusCode)
			}
			fmt.Printf("%d\n", i)
		}(i)
	}

	wg.Wait()
}

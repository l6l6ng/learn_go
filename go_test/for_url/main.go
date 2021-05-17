package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	params := url.Values{}

	for i := 100; i <= 102; i++ {
		url1 := "http://192.168.1.37:8099/api/tranform"
		Url, err := url.Parse(url1)
		if err != nil {
			panic(err.Error())
		}
		//params.Set("url", "http://192.168.2.202/index.php/Bashu/Report/310997")
		params.Set("url", "http://192.168.2.202/index.php/Report/reportTerm2017/310997")
		str := "/学院/学院" + strconv.Itoa(i) + ".pdf"
		fmt.Println(i)
		fmt.Println(str)
		params.Set("targetPath", string(str))
		params.Set("callBack", "https://www.baidu.com")
		//url = "https://api.github.com/users/dfsd534"
		Url.RawQuery = params.Encode()
		urlPath := Url.String()
		fmt.Println(urlPath)
		fmt.Println("/学院/学院" + string(i) + ".pdf")
		urlPath = "https://www.baidu.com"
		resp, err := http.Get(urlPath)
		if err == nil {
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				//var b string
				fmt.Println(string(body))
				//err := json.Unmarshal(body,&b)
				//if err == nil {
				//	fmt.Println(b)
				//} else {
				//	fmt.Println(err)
				//}
			} else {
				fmt.Println("e1", err)
			}
		} else {
			fmt.Println("e2", err)
		}
		resp.Body.Close()
	}

}

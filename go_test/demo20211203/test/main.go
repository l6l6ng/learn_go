package main

import (
	"learn_go/demo20211203/mylog"
)

func main() {
	f1 := mylog.NewFileLogger(mylog.INFO, "./", "test.log")
	userId := 100
	f1.Debug("%v记录了一条日志", userId)
}

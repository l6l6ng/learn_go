package mylog

import (
	"path"
	"runtime"
)

func getCallerInfo() (file, funcName string, line int) {
	pc, file, line, ok := runtime.Caller(2)
	funcName = runtime.FuncForPC(pc).Name()
	if !ok {
		return
	}

	funcName = path.Base(funcName)
	file = path.Base(file)

	return
}

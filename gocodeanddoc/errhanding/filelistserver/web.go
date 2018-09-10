package main

import (
	"mygolearn/gocodeanddoc/errhanding/filelistserver/filelist"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

//输入函数包装为输出函数取使用
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			//此函数有三个参数，一个是writer，一个是字符串code，一个是code
			http.Error(writer,
				http.StatusText(code),
				code)
		}
	}
}

func main() {
	//对应地址的处理函数
	http.HandleFunc("/list/",
		errWrapper(filelist.HandlerFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"learngo/07errorhandling/filelisting"
	// "learngo/07errorhandling/recovery"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

type userError interface {
	error
	Message() string
}

// 函数式编程，输入输出都是函数
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := filelisting.HandleFileList(writer, request)
		if err != nil {
			log.Println("Error handling request: ", err.Error())
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func tryErrorHandling() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	tryErrorHandling()
	// recovery.TryRecover()
}

// 意料之中的，用error，如：文件打不开
// 意料之外的，用panic，如：数组越界

package main

import (
	"bufio"
	"fmt"
	"learngo/07error/fib"
	"os"
)

func tryDefer() {
	// defer确保调用在函数结束时发生
	// 参数在defer语句时计算
	// defer列表为先进后出
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occured")
	return
	fmt.Println(4)
}

func writeFile(filename string) {
	// 常用情景：open/close  lock/unlock  printheader/printfooter
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Fprintln(writer, f())
	}
}

func writeFile2(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()
	writeFile("fib.txt")
	// writeFile2("fib.txt")
}

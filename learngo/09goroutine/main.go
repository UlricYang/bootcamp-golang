package main

import (
	"fmt"
	"runtime"
	"time"
)

const N = 8

func main() {
	var a [N]int
	for i := 0; i < N; i++ {
		go func(j int) {
			for {
				a[j]++
				runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Millisecond)

	aa := 0
	for _, item := range a {
		aa += item
	}
	fmt.Println(a)
	fmt.Println(aa)
}

// goroutine 可能切换的点:
// IO,select
// channel
// runtime.Gosched
// 等待锁
// 函数调用（有时）

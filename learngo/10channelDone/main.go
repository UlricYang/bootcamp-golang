package main

import (
	"fmt"
	"sync"
)

const N = 26

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %2d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int, 0),
		done: func() { wg.Done() },
	}
	go doWork(id, w)
	return w
}

func channelDemo() {
	var wg sync.WaitGroup

	var workers [N]worker
	for i := 0; i < N; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(len(workers) * 2)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

// func bufferedChannel() {
// 	c := make(chan int, 3)
// 	go worker(0, c)
// 	c <- 1
// 	c <- 2
// 	c <- 3
// 	c <- 4
// 	time.Sleep(time.Millisecond)
//
// }

// func channelClose() {
// 	c := make(chan int, 3)
// 	go worker(0, c)
// 	c <- 1
// 	c <- 2
// 	c <- 3
// 	c <- 4
// 	close(c)
// 	time.Sleep(time.Millisecond)
// }

func main() {
	channelDemo()
}

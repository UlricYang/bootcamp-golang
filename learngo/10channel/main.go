package main

import (
	"fmt"
	"time"
)

const N = 26

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %2d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %2d received %c\n", id, <-c)
		}
	}()
	return c
}

func channelDemo() {
	var channels [N]chan<- int
	for i := 0; i < N; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < N; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < N; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)

}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
	fmt.Println()
	bufferedChannel()
	fmt.Println()
	channelClose()
}

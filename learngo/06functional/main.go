package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

type intGen func() int

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}

	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a2 = a2(i)
		fmt.Println(i, s, a2)

	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("fib(%d)=%d\n", i, f())
	}

	g := fibonacci()
	printFileContents(g)
}

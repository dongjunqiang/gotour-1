package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		fmt.Println("fib", i, x)
		x, y = y, x+y
	}
	close(c)
}

func main() {
	//c := make(chan int, 23)
	//c := make(chan int, 42)
	c := make(chan int, 64)
	go fibonacci(cap(c), c)
	//fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}

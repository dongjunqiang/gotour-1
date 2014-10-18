package main

import "fmt"

func main() {
	c := make(chan int)
	fmt.Println("cap", cap(c), "len", len(c))
	ch := make(chan int, 64)
	fmt.Println("cap", cap(ch), "len", len(ch))
	for i := 0; i < 42; i++ {
		ch <- i
	}
	fmt.Println("cap", cap(ch), "len", len(ch))
}

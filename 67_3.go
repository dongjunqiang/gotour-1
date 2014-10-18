package main

import "fmt"

func main() {
	var c chan int
	fmt.Println("cap", cap(c), "len", len(c))
}

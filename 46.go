package main

import "fmt"

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		t := x
		x, y = y, x+y
		return t
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

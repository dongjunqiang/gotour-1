package main

import "fmt"

func f(i int) {
	for {
		fmt.Println("goroutine", i)
	}
}

func main() {
	for i := 0; i < 23; i++ {
		go f(i)
	}

	for {
		fmt.Println("main")
	}
}

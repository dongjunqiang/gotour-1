package main

import (
	"fmt"
	"runtime"
	"time"
)

func Sleep(seconds int) {
	for i := 0; i < seconds; i++ {
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))
	fmt.Println("GOROOT", runtime.GOROOT())
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("NumGoroutine", runtime.NumGoroutine())

	for i := 0; i < 64; i++ {
		go Sleep(i)
	}

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("NumGoroutine", runtime.NumGoroutine())
	}
}

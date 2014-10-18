package main

import "fmt"
//import "time"

func f() {
    fmt.Println("goroutine")
}

func main() {
    go f()

    for {
        //time.Sleep(1 * time.Second)
    }
}

package main

import "fmt"

func Foo() (x int, y int) {
	return
}

func Bar() (x, y int) {
	x = 42
	return
}

func Foobar() (x, y int) {
	x, y = 42, 23
	return
}

func main() {
	fmt.Println(Foo())
	fmt.Println(Bar())
	fmt.Println(Foobar())
}

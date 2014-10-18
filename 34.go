package main

import "fmt"

func PrintSlice(s string, x []int) {
	fmt.Printf("%s %T len=%d cap=%d %v\n",
		s, x, len(x), cap(x), x)
}

func main() {
	a := make([]int, 5)
	PrintSlice("a", a)
	b := make([]int, 0, 5)
	PrintSlice("b", b)
	c := b[:2]
	PrintSlice("c", c)
	d := c[2:5]
	PrintSlice("d", d)
}

package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println()
	for i := 1; i <= 9; i++ {
		fmt.Println(i)
	}

	fmt.Println()
	for {
		fmt.Println("loop")
		break
	}
}

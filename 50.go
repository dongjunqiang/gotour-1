package main

import "fmt"

func Cbrt(x complex128) complex128 {
	z := complex(1.0, 0)
	for i := 0; i < 64; i++ {
		z = z - (z*z*z-x)/(3*z*z)
	}
	return z
}

func main() {
	for i := 0; i < 64; i++ {
		fmt.Println(i, Cbrt(complex(float64(i), 0)))
	}
}

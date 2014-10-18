package main

import "strings"
import "code.google.com/p/go-tour/wc"

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	fields := strings.Fields(s)
	for _, value := range fields {
		m[value] = m[value] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

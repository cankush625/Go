package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for _, val := range a {
		m[val]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

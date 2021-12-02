package main

import "fmt"

func ptr() {
	var a *int
	b := 10
	fmt.Println(a) // output: <nil>
	fmt.Println(b) // output: 10
	a = &b
	fmt.Println(a) // output: memory address of a value a
	fmt.Println(*a) // output: 10
	fmt.Println(b) // output: 10
}

func main() {
	ptr()
}

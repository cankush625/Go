package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello")
	}()

	val := func() int {
		return 5
	}()
	fmt.Println(val)
}

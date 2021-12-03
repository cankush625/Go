package main

import "fmt"

func addFive() func(int) int {
	five := 5
	return func(num int) int {
		sum := num + five
		return sum
	}
}

func main() {
	sumFunction := addFive()
	fmt.Println(sumFunction(2))
}

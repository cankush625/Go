package main

import (
	"fmt"
)

func addFirstTenNumbers(startNum int) int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += startNum + i
	}
	return sum
}

func isEven(num int) bool {
	if num % 2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	num := 10
	result := isEven(num)
	fmt.Printf("Is %v even? %v\n", num, result)
	startNum := 1
	sum := addFirstTenNumbers(startNum)
	fmt.Printf("The sum of numbers from %v to %v is %v", startNum, startNum + 9, sum)
}

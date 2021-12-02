package main

import "fmt"

func main() {
	var arr [10]int
	arr[0] = 2
	arr[6] = 5
	fmt.Println(arr)

	primes := [5]int{2, 3, 5}
	fmt.Println(primes)
}

package main

import "fmt"

// add the sum of elements to the passed channel
func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func main() {
	channel := make(chan int)
	arr := []int{1, 2, 3, 4, 5}
	go sum(arr[:2], channel)
	go sum(arr[2:], channel)
	x, y := <-channel, <-channel
	fmt.Println(x, y, x+y)

	// buffered channel
	// Provide buffer length as the second argument
	ch := make(chan int, 2)
	ch <- 1
	ch <- 5
	fmt.Println(<-ch, <-ch)
}

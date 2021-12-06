package main

import "fmt"

func storeValue(capacity int, channel chan int) {
	for i := 0; i < capacity; i++ {
		channel <- i + 1
	}
	close(channel)
}

func main() {
	channel := make(chan int, 2)
	go storeValue(cap(channel), channel)
	for i := range channel {
		fmt.Println(i)
	}
}

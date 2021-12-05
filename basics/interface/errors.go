package main

import "fmt"

type MyError struct {
	name string
	message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.name, e.message)
}

func run() error {
	return MyError{
		"KeyError",
		"Key not found",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

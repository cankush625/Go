package main

import "fmt"

type Person1 struct {
	Name string
	Age  int
}

// The fmt package (and many others) look for this interface to print values.
func (p Person1) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person1{"Bob Smith", 42}
	z := Person1{"abc sshd", 67}
	fmt.Println(a, z)
}

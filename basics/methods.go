package main

import "fmt"

type Car struct {
	name  string
	brand string
	price int
}

// getter method using value receiver
func (c Car) getCarName() string {
	return c.name
}

// setter method using pointer receiver
func (c *Car) setCarPrice(price int) {
	c.price = price
}

func main() {
	db11 := Car{"db11", "Aston Martin", 1000}
	carName := db11.getCarName()
	fmt.Println(carName)
	fmt.Println(db11.price)
	db11.setCarPrice(1500)
	fmt.Println(db11.price)
}

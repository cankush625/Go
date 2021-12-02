package main

import "fmt"

type Person struct {
	name string
	age  int
}

func structLiterals() {
	var (
		person1 = Person{"Ricky", 65}  // has type Person
		person2 = Person{name: "Mark"} // age:0 is implicit
		person3 = Person{}             // name:"" and age:0
		person4 = &Person{"David", 67} // has type *Person
	)
	fmt.Println(person1, person2, person3, person4)
}

func main() {
	fmt.Println(Person{"ab", 24})
	person := Person{"Rob", 32}
	fmt.Println(person.name)
	fmt.Println(person.age)
	person.age = 34
	fmt.Println(person.name)
	fmt.Println(person.age)

	// Pointers to structs
	ptr := &person
	ptr.name = "James"
	fmt.Println(person.name)
	fmt.Println(person.age)

	// Struct literals
	structLiterals()
}

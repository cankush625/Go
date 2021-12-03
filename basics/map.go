package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func main() {
	// Create map using make function
	m := make(map[string]Animal)
	m["dog"] = Animal{
		"dog",
		10,
	}
	fmt.Println(m)
	fmt.Println(m["dog"])

	// Create map directly without using make function
	// Map literals
	var s = map[int]Animal{0: {"tiger", 2}}
	fmt.Println(s)
	fmt.Println(s[0])

	// Mutating Maps
	// Insert or update an element
	s[0] = Animal{
		"tiger",
		20,
	}
	s[1] = Animal{
		"lion",
		25,
	}
	fmt.Println(s)
	fmt.Println(s[0])

	// Delete an element
	delete(s, 1)
	fmt.Println(s)

	// Test if key is present in the map
	_, ok := s[1]
	fmt.Println("The key", 1, "present?", ok)
}

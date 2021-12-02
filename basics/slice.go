package main

import "fmt"

func main() {
	primes := [5]int{2, 3, 5, 7, 11}

	var sliced []int = primes[1:4]
	fmt.Println(sliced)

	var firstToSpecified = primes[:3]
	fmt.Println(firstToSpecified)
	firstToSpecified = primes[0:3]
	fmt.Println(firstToSpecified)
	fmt.Println(
		"Length", len(firstToSpecified), "and capacity",
		cap(firstToSpecified), "of firstToSpecified",
	)

	var specifiedToLast = primes[2:]
	fmt.Println(specifiedToLast)
	fmt.Println(
		"Length", len(specifiedToLast), "and capacity",
		cap(specifiedToLast), "of firstToSpecified",
	)

	var includeAll = primes[:]
	fmt.Println(includeAll)

	// Slice literals
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// Slice of team
	team := []struct {
		name   string
		points int
	}{
		{"Sachin Tendulkar", 98},
		{"MS Dhoni", 95},
		{"AB Deviliars", 96},
	}
	fmt.Println(team)
}

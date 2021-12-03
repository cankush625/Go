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

	// Slice with make
	sl := make([]int, 5)
	fmt.Println(sl)
	fmt.Println("Length", len(sl), "and capacity", cap(sl), "of sl")

	// Slice with make and specify capacity using third argument to make
	// Slice of length 2 and capacity 5
	sls := make([]int, 2, 5)
	fmt.Println(sls)
	fmt.Println("Length", len(sls), "and capacity", cap(sls), "of sls")

	// Slices of slices
	board := [][]string{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
	}
	fmt.Println(board)
	board[1][2] = "g"
	fmt.Println(board)

	// Append to slice
	s := []int{1, 2}
	fmt.Println(s)
	s = append(s, 3)
	fmt.Println(s)
	// append multiple values
	s = append(s, 4, 5)
	fmt.Println(s)

	// range with slice
	square := []int{0, 1, 4, 9, 16, 25}
	for i, sq := range square {
		fmt.Printf("Square of %d is %d\n", i, sq)
	}

	// get only index from slice
	for i := range square {
		fmt.Println(i)
	}

	// skip the index and get only value from range with slice
	for _, val := range square {
		fmt.Println(val)
	}
}

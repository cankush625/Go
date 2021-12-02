package main

import (
	"fmt"
	"os"
)

func printError() {
	const name, id = "as", 673
	err := fmt.Errorf("user %q with id %d is not found", name, id)
	fmt.Println(err.Error())
}

func usingFprint() {
	const name, age = "Smith", 25
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")

	// Print the error if any
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")
}

func usingFprintf() {
	const name, age = "Smith", 25
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// Print the error if any
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)
}

func usingFprintln() {
	const name, age = "Smith", 25
	// Space will always be added between operands and a newline is appended
	n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

	// Print the error if any
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Println(n, "bytes written.")
}

func usingPrint() {
	const name, age = "Smith", 25
	n, err := fmt.Print(name, " is ", age, " years old.\n")

	// Print the error if any
	if err != nil {
		_, _ = fmt.Printf("Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")
}

func usingPrintf() {
	const name, age = "Smith", 25
	n, err := fmt.Printf("%s is %d years old.\n", name, age)

	// Print the error if any
	if err != nil {
		_, _ = fmt.Printf("Fprint: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)
}

func usingPrintln() {
	const name, age = "Smith", 25
	// Space will always be added between operands and a newline is appended
	n, err := fmt.Println(name, "is", age, "years old.")

	// Print the error if any
	if err != nil {
		_, _ = fmt.Printf("Fprint: %v\n", err)
	}
	fmt.Println(n, "bytes written.")
}

func usingScanln() {
	var name string
	fmt.Print("Enter your name: ")
	n, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("My name is", name)
	fmt.Printf("%d bytes written.\n", n)
}

func main() {
	printError()
	usingFprint()
	usingFprintf()
	usingFprintln()
	usingPrint()
	usingPrintf()
	usingPrintln()
	usingScanln()
}

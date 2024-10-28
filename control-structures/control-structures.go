package main

import "fmt"

func main() {
	fmt.Println("Control structures in Go")

	// If, else if, else
	fmt.Println("If, else if, else in Go")
	num := 10
	if num%2 == 0 {
		fmt.Printf("%d is odd\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	if num < 0 {
		fmt.Printf("%d is negative\n", num)
	} else if num > 0 {
		fmt.Printf("%d is positive\n", num)
	} else {
		fmt.Printf("%d is zero\n", num)
	}

	// Switch
	fmt.Println("Switch in Go")
	num = 2
	switch num {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Default")
	}
}

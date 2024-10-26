package main

import "fmt"

func main() {
	fmt.Println("Data structures in Go")

	// Arrays
	fmt.Println("Arrays")
	var myArray [5]int
	fmt.Println(myArray)
	myArray[0] = 1234
	fmt.Println(myArray)

	// Slices
	fmt.Println("Slices")
	var mySlice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)
	mySlice = append(mySlice, 6)
	fmt.Println(mySlice)

	// Maps
	fmt.Println("Maps")
	var myMap map[string]int = map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap)
	fmt.Println(myMap["a"])
	myMap["c"] = 3
	fmt.Println(myMap)
	delete(myMap, "a")
	fmt.Println(myMap)
}

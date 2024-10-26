package main

import "fmt"

func sumTuple(a int, b int) (int, int) {
	return a + b, b + a
}

func main() {
	fmt.Println("Variables and types in Go")
	fmt.Println("Go" + "Lang")
	fmt.Println(false || true)
	fmt.Println(!false)
	fmt.Println(false && true)
	var integer int = 1
	var str string = "String"
	fmt.Println(integer)
	fmt.Println(str)
	var a, b int = sumTuple(1, 2)
	fmt.Println(a, b)
}

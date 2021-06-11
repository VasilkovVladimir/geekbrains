package main

import "fmt"

func main() {
	var a, b float64
	fmt.Print("First side length: ")
	fmt.Scanln(&a)
	fmt.Print("Second side length: ")
	fmt.Scanln(&b)
	s := a * b
	fmt.Println("Rectangle area: ", s)
}
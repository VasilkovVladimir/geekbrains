package main

import (
	"fmt"
)

func main()  {
	var a int16
	fmt.Print("Enter tree-digit number:")
	fmt.Scanln(&a)
	fmt.Println("Number of hundreds:", a / 100)
	fmt.Println("Number of tens:", (a % 100) / 10)
	fmt.Println("Number of ones:", a % 10)
}
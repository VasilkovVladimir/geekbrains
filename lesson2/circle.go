package main

import (
	"fmt"
	"math"
)

func main()  {
	var s, d, l, r float64
	fmt.Print("Enter circle area:")
	fmt.Scanln(&s)
	r = math.Sqrt(s / math.Pi)
	d = r * 2
	l = d * math.Pi
	fmt.Println("Diameter is:", d)
	fmt.Println("Circle length:", l)
}
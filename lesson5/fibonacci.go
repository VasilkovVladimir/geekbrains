package main

import "fmt"

func Fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibo(n-1) + Fibo(n-2)
	}
}

func main() {
	var x int
	results := map[int]int{
		0: 0,
		1: 1,
	}
	fmt.Print("Введите номер числа в последовательности: ")
	fmt.Scanln(&x)
	fmt.Println(Fibo(x))
	results[x] = Fibo(x)
	fmt.Println(results)
}

/*if x,ok := result[x] {
	return
} else {
	result[x] = Fibo(x)
}*/
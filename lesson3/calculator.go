package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("Не является числом")
		return
	}

	fmt.Print("Введите второе число: ")
	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Println("Не является числом")
		return
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, %): ")
	if _, err := fmt.Scanln(&op); err != nil {
		return
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "%":
		res = math.Mod(a, b)

	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

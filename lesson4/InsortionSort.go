package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for ; j >= 1 && arr[j-1] > x; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = x
	}
}

func main() {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	InsertionSort(arr)
	fmt.Println(arr)
}

package main

import "fmt"

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T

	for _, v := range slice {
		sum += v
	}

	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func Example01() {
    n := []int{1, 2, 3}
	fmt.Println(sumSlice(n))
	fmt.Println(sumSlice([]float32{1, 2, 3}))
	fmt.Println(sumSlice[float64]([]float64{1, 2, 3}))
	// fmt.Println(sumSlice([]uint{1, 2, 3})) // error

    fmt.Println(isEmpty(n))
    fmt.Println(isEmpty([]int{}))
}

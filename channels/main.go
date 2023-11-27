package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		sum := 0
		for i := 0; i < 100; i++ {
			fmt.Println("IDX from first func:", i)
			sum += i
		}
        c <- sum
	}()

    output := <-c
    fmt.Println("Output:", output)
}

package main

import (
	"fmt"
	"time"
)

func Take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)

	go func() {
		defer close(taken)
		
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()

	return taken
}

func PrimeFinder[K any](done <-chan K, randIntStream <-chan int) <-chan int {
	// creating a slow pipeline stage
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt % i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomInt := <- randIntStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()

	return primes
}

func Example02() {
	start := time.Now()
	done := make(chan struct{})
	defer close(done)

	randStream := RepeatFunc(done, RandomNumFetcher)
	primeStream := PrimeFinder(done, randStream)

	// naive/slow solution
	for rando := range Take(done, primeStream, 5) {
		fmt.Println(rando)
	}

	fmt.Println(time.Since(start))
}

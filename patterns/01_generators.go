package main

import (
	"fmt"
)

func RepeatFunc[K any, T any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func Example01() {
	done := make(chan struct{})
	defer close(done)

	for rando := range RepeatFunc(done, RandomNumFetcher) {
		fmt.Println(rando)
	}
}

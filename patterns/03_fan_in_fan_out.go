package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func fanOut[K any](done <-chan K, randIntStream <-chan int) []<-chan int {
	CPUCount := runtime.NumCPU()
	finderChannels := make([]<-chan int, CPUCount)

	for i := 0; i < CPUCount; i++ {
		finderChannels[i] = PrimeFinder(done, randIntStream)
	}

	return finderChannels
}


func fanIn[K any, T any](done <-chan K, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}

func Example03() {
	start := time.Now()

	done := make(chan struct{})
	defer close(done)

	randStream := RepeatFunc(done, RandomNumFetcher)
	primeFinderChannels := fanOut(done, randStream)
	fannedPrimesStream := fanIn(done, primeFinderChannels...)

	for result := range Take(done, fannedPrimesStream, 10) {
		fmt.Println(result)
	}

	fmt.Println(time.Since(start))
}

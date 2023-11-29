package main

import (
	"fmt"
	"sync"
	"time"
)

func countDown(id string, n int) {
	for n >= 0 {
		fmt.Println(id, n)
		n--
		time.Sleep(time.Millisecond * 500)
	}
}

func Example03() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		countDown(">>", 5)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		countDown("#", 10)
	}()

	wg.Wait()
}

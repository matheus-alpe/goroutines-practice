package main

import (
	"fmt"
	"time"
)

func Worker(done chan struct{}) {
	fmt.Println("Working")
	time.Sleep(time.Second * 2)
	fmt.Println("Done")
	close(done)
}

func Example05() {
	fmt.Println("Starting workers")

	done := make(chan struct{})
	go Worker(done)
	<-done

	fmt.Println("Ending workers")
}

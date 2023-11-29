package main

import (
	"fmt"
	"time"
)

func Example06() {
	fmt.Println("Starting workers")

	ch := make(chan int, 2)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")
			time.Sleep(time.Second)
		}

		fmt.Println(time.Now(), "all completed, leaving")
		close(ch)
	}()

	go func() {
		for v := range ch {
			fmt.Println(time.Now(), "received", v)
		}
		close(done)
		// for {
		// 	select {
		// 	case v, open := <-ch:
		// 		if !open {
		// 			close(done)
		// 			return
		// 		}
		// 	default:
		// 		continue
		// 	}
		// }
	}()
	<-done

	fmt.Println("Ending workers")
}

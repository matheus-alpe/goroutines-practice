package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "123456")
}

func doSomething(ctx context.Context) {
	rId := ctx.Value("request-id")

outter:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			break outter
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(time.Second * 1)
	}

	fmt.Println(rId)
}

func Example01() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	ctx = enrichContext(ctx)
	go doSomething(ctx)
	<-ctx.Done() // will block here until is timed out
	fmt.Println(ctx.Err())
	time.Sleep(time.Second * 2)
}

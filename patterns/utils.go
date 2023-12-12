package main

import "math/rand"

func RandomNumFetcher() int {
	return rand.Intn(5e8)
}

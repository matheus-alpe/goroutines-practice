package main

import (
	"log"
	"os"
)

func main() {
    if err := root(os.Args[1:]); err != nil {
        log.Fatalln(err)
    }
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_PC_PRICE float32 = 5
const MAX_KB_PRICE float32 = 5

func checkPcPrices(website string, pcChan chan string) {
	for {
		time.Sleep(time.Second)
		randomPrice := rand.Float32() * 20
		
		if randomPrice <= MAX_PC_PRICE {
			pcChan <- website
			fmt.Println(randomPrice)
			break
		}
	}
}
func checkKeyboardPrices(website string, kbChan chan string) {
	for {
		time.Sleep(time.Second)
		randomPrice := rand.Float32() * 20
		
		if randomPrice <= MAX_PC_PRICE {
			kbChan <- website
			fmt.Println(randomPrice)
			break
		}
	}
}

func sendMessage(pcChan chan string, kbChan chan string) {
	select {
	case pcWebsite := <-pcChan:
		fmt.Println("PC deal founded at", pcWebsite)
	case kbWebsite := <-kbChan:
		fmt.Println("Keyboard deal founded at", kbWebsite)
	}

	// fmt.Println("PC deal founded at", <-pcChan)
	// fmt.Println("Keyboard deal founded at", <-kbChan)
}

func Example07() {
	pcChan := make(chan string)
	kbChan := make(chan string)
	websites := [...]string{"walmart.com", "amazon.com", "kabum.com.br"}
	
	for i := range websites {
		go checkPcPrices(websites[i], pcChan)
		go checkKeyboardPrices(websites[i], kbChan)
	}

	sendMessage(pcChan, kbChan)
}

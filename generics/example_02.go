package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	name  string
	email string
}

type purchaseInfo struct {
	name   string
	price  float32
	amount int
}

func loadJSON[T contactInfo | purchaseInfo](filepath string) []T {
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}

func Example02() {
	contacts := loadJSON[contactInfo]("./generics/contactInfo.json")
	fmt.Printf("\n%+v", contacts)

	purchases := loadJSON[purchaseInfo]("./generics/purchaseInfo.json")
	fmt.Printf("\n%+v", purchases)
}

package main

import (
	"fmt"
)

type GasEngine struct {
	gallons float32
	mpg float32
}

type EletricEngine struct {
	kwh float32
	mpkwh float32
}

type Car [T GasEngine | EletricEngine] struct {
	make, model string
	engine T
}

func Example03() {
	gasCar := Car[GasEngine]{
		make: "Honda",
		model: "Civic",
		engine: GasEngine{
			gallons: 12.4,
			mpg: 40,
		},
	}

	eletricCar := Car[EletricEngine]{
		make: "Tesla",
		model: "Model 3",
		engine: EletricEngine{
			kwh: 57.5,
			mpkwh: 4.17,
		},
	}

	fmt.Println(gasCar)
	fmt.Println(eletricCar)
}

package main

import (
	"fmt"
	"popashopa/feature_1"
)

type Auto struct {
	Name string
	Power int
}

func main()  {
	
	popa := []Auto{{Name: "bmv", Power: 1234}, {Name: "Mercades", Power: 123}}
	for _, av := range popa {
		fmt.Printf("Автомобиль: %10s | Мощность: %4d\n", av.Name, av.Power)
	}

	popa[0].Name = "pizda"
	popa[0].Power = 1235

	for _, av := range popa {
		fmt.Printf("Автомобиль: %10s | Мощность: %4d\n", av.Name, av.Power)
	}

	feature1.Feature1()

}
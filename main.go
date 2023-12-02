package main

import (
	"fmt"
	"log"

	"bartekpacia/bit-festival-2023/algo"
)

func main() {
	err := algo.Init()
	if err != nil {
		log.Fatalln(err)
	}

	I_obl := 18.0 // A
	temp := 27.0  // °C
	fmt.Printf("I_obl: %f A, temp: %f °C\n", I_obl, temp)
	I_ost, err := algo.CalcTemp(I_obl, temp)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("I_ost:", I_ost)
}

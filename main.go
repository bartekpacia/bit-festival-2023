package main

import (
	"fmt"
	"log"

	"bartekpacia/bit-festival-2023/algo"
)

func main() {
	algo.LoadPath = "algo/"
	err := algo.Init()
	if err != nil {
		log.Fatalln(err)
	}

	I_obl := 18.0 // A
	temp := 27.0  // °C
	P := 0.0
	fmt.Printf("I_obl: %f A, temp: %f °C\n", I_obl, temp)
	I_ost, err := algo.CalcTemp(I_obl, P, temp, 3)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("I_ost:", I_ost)

	przekroj, err := algo.MatchCrossection(I_ost, 3, algo.A1)
	if err != nil {
		log.Fatalln("failed to match crossection:", err)
	}

	fmt.Println("przekroj:", przekroj)
}

package main

import (
	"fmt"

	"bartekpacia/bit-festival-2023/algo"
)

func main() {
	I_obl := 18.0 // A
	Temp := 0.94  // współczynnik tolerancji dla 35°C dla kabla w powietrzu
	got_I_ost := algo.Calc(
		I_obl, /* maksymalne obciążenie kabla */
		Temp,  /*temperatura otoczenia*/
	)

	fmt.Println(got_I_ost)
}

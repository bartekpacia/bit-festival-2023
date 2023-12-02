package algo

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 0.1

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestCalc_1(t *testing.T) {
	// CASE 1 ładowarka 18A na saharze w powietrzu 35°C

	I_obl := 18.0 // A
	Temp := 0.94  // współczynnik tolerancji dla 35°C dla kabla w powietrzu
	got_I_ost := Calc(
		I_obl, /* maksymalne obciążenie kabla */
		Temp,  /*temperatura otoczenia*/
	)

	const want_I_ost = 22.53 // tyle wyszło z obliczeń ręcznych
	if !almostEqual(got_I_ost, want_I_ost) {
		t.Errorf("Calc(%f, %f) = %f; want %f", I_obl, Temp, got_I_ost, want_I_ost)
	}
}

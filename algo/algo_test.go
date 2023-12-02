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
	k := 0.94     // współczynnik tolerancji dla 35°C dla kabla w powietrzu
	got_I_ost := Calc(
		I_obl, /* maksymalne obciążenie kabla */
		k,     /*temperatura otoczenia*/
	)

	const want_I_ost = 22.53 // tyle wyszło z obliczeń ręcznych
	if !almostEqual(got_I_ost, want_I_ost) {
		t.Errorf("Calc(%f, %f) = %f; want %f", I_obl, k, got_I_ost, want_I_ost)
	}
}

func TestCalcTemp_1(t *testing.T) {
	// CASE 1 ładowarka 18A na saharze w powietrzu 27°C

	I_obl := 18.0 // A, maksymalne obciążenie prądowe kabla
	temp := 27.0  // 27°C, temperatura otoczenia
	got_I_ost, _ := CalcTemp(I_obl, temp)

	const want_I_ost = 21.1765 // tyle wyszło z obliczeń ręcznych
	if !almostEqual(got_I_ost, want_I_ost) {
		t.Errorf("CalcTemp(%f, %f) = %f; want %f", I_obl, temp, got_I_ost, want_I_ost)
	}
}

package algo_test

import (
	"log"
	"math"
	"os"
	"testing"

	"bartekpacia/bit-festival-2023/algo"
)

const float64EqualityThreshold = 0.1

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestMain(m *testing.M) {
	err := algo.Init()
	if err != nil {
		log.Println(os.Getwd())
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}

func TestCalc_1(t *testing.T) {
	// CASE 1 ładowarka 18A na saharze w powietrzu 35°C

	I_obl := 18.0 // A (maksymalne obciążenie prądowe kabla)
	k := 0.94     // współczynnik tolerancji dla 35°C dla kabla w powietrzu
	got_I_ost := algo.Calc(I_obl, k)

	want_I_ost := 22.53 // tyle wyszło z obliczeń ręcznych
	if !almostEqual(got_I_ost, want_I_ost) {
		t.Errorf("Calc(%f, %f) = %f; want %f", I_obl, k, got_I_ost, want_I_ost)
	}
}

func TestCalcTemp_1(t *testing.T) {
	I_obl := 18.0 // A (maksymalne obciążenie prądowe kabla)
	temp := 5.0   // °C (temperatura otoczenia)
	got_I_ost, _ := algo.CalcTemp(I_obl, temp)

	want_I_ost := 17.3578 // tyle wyszło z obliczeń ręcznych
	if !almostEqual(got_I_ost, want_I_ost) {
		t.Errorf("CalcTemp(%f, %f) = %f; want %f", I_obl, temp, got_I_ost, want_I_ost)
	}
}

func TestCalcTemp_2(t *testing.T) {
	I_obl := 18.0 // A (maksymalne obciążenie prądowe kabla)
	temp := 27.0  // °C (temperatura otoczenia)
	got_I_ost, _ := algo.CalcTemp(I_obl, temp)

	want_I_ost := 21.1765 // tyle wyszło z obliczeń ręcznych
	if !almostEqual(got_I_ost, want_I_ost) {
		t.Errorf("CalcTemp(%f, %f) = %f; want %f", I_obl, temp, got_I_ost, want_I_ost)
	}
}

func TestCalcTemp_3(t *testing.T) {
	I_obl := 18.0 // A (maksymalne obciążenie prądowe kabla)
	temp := 57.0  // °C (temperatura otoczenia)
	got_I_ost, _ := algo.CalcTemp(I_obl, temp)

	want_I_ost := 42.3529 // tyle wyszło z obliczeń ręcznych
	if !almostEqual(got_I_ost, want_I_ost) {
		t.Errorf("CalcTemp(%f, %f) = %f; want %f", I_obl, temp, got_I_ost, want_I_ost)
	}
}

func TestMatchCrossection_1(t *testing.T) {
	I_ost := 22.53 // A (prąd obciążenia)
	zyly := 2      // liczba przewodów
	location := algo.A1
	got, err := algo.MatchCrossection(I_ost, zyly, location)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	want := 4.0 // mm2 (przekrój przewodu)
	if !almostEqual(got, want) {
		t.Errorf("MatchCrossection(%f, %d, %s) = %f; want %f", I_ost, zyly, location, got, want)
	}
}

func TestMatchCrossection_2(t *testing.T) {
	I_ost := 90.25 // A (prąd obciążenia)
	zyly := 3      // liczba przewodów
	location := algo.A2
	got, err := algo.MatchCrossection(I_ost, zyly, location)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	want := 50.0 // mm2 (przekrój przewodu)
	if !almostEqual(got, want) {
		t.Errorf("MatchCrossection(%f, %d, %s) = %f; want %f", I_ost, zyly, location, got, want)
	}
}

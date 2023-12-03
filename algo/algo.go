package algo

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	ErrNotInitialized = fmt.Errorf("algo package not initialized")
	ErrNotFound       = fmt.Errorf("not found")
)

var (
	LoadPath      string
	YDY_tolerance [][]string
	YDY_przekroj  [][]string
)

// Init loads data needed for the algorithm to work.
func Init() error {
	{
		file, err := os.Open(fmt.Sprintf("%stolerance_data.csv", LoadPath))
		if err != nil {
			return fmt.Errorf("failed to load tolerance csv file: %w", err)
		}
		defer file.Close()

		// read csv values using csv.Reader
		csvReader := csv.NewReader(file)
		YDY_tolerance, err = csvReader.ReadAll()
		if err != nil {
			return fmt.Errorf("failed to read from file into csv: %w", err)
		}
	}

	{
		file, err := os.Open(fmt.Sprintf("%sYDY_przekroj.csv", LoadPath))
		if err != nil {
			return fmt.Errorf("failed to load przekroj csv file: %w", err)
		}
		defer file.Close()

		// read csv values using csv.Reader
		csvReader := csv.NewReader(file)
		YDY_przekroj, err = csvReader.ReadAll()
		if err != nil {
			return fmt.Errorf("failed to read from file into csv: %w", err)
		}
	}

	return nil
}

func Calc(I_obl, tolerance float64) (I_ost float64) {
	I_ost = I_obl / (tolerance * 0.85)
	return
}

// CalcTemp returns I_ost for given I_obl and temp.
func CalcTemp(I_obl, P, temp float64, zyly int) (float64, error) {
	if YDY_tolerance == nil {
		panic(ErrNotInitialized)
	}

	// find to which cable we should match
	przypadek := 1
	if zyly == 3 {
		zyly = 2
		przypadek = 1 //jako przypadek wielożyłowego kabla
	} else if zyly == 4 || zyly == 5 {
		zyly = 3
		przypadek = 1 //jako przypadek wielożyłowego kabla
	} else if zyly == 1 {
		zyly = 3
		przypadek = 2 //jako przypadek jednożyłowego kabla
	}

	if przypadek == 1 {
		//przeskanuj i wyświetl wynik tylko dla kabli YDY, YDYp, YKY,
	} else if przypadek == 2 {
		//przeskanuj i wyświetl wynik tylko dla kabli YKY, YKXS, YAKXS, N2XH
	}

	if P != 0.0 {
		if zyly == 2 { //kabel 1 fazowy
			I_obl = P * 1000 / (230 * 0.8)
		} else if zyly == 3 { //kabel 3 fazowy
			I_obl = P * 1000 / (math.Sqrt(3) * 400 * 0.8)
		}
	}
	// start from 1 to skip first column
	for i := 1; i < len(YDY_tolerance); i++ {
		thisRow := YDY_tolerance[i]
		thisTemp, _ := strconv.ParseFloat(thisRow[0], 64)
		nextTolerance, _ := strconv.ParseFloat(thisRow[1], 64)
		if thisTemp > temp {
			return Calc(I_obl, nextTolerance), nil
		}
	}

	return 0, ErrNotFound
}

type CableLocation string

const (
	// przewód wielożyłowy bezpośrednio w ścianie izolowanej cieplnie
	A1 CableLocation = "A1"
	// przewód wielożyłowy w rurze instalacyjnej w ścianie izolowanej cieplnie
	A2 CableLocation = "A2"
	// przewód  w rurze instalacyjnej/listwie naściennej na ścianie murowanej
	// lub drewnianej
	B2 CableLocation = "B2"
	// przewód wielożyłowy w powietrzu (korytko perforowane, siatkowe)
	E CableLocation = "E2"
)

// MatchCrossection zwraca przekrój przewodu dla podanego I_ost, zyly i gdzie.
// I_ost = ostateczny prąd obciążenia
func MatchCrossection(I_ost float64, zyly int, gdzie CableLocation) (float64, error) {
	if YDY_przekroj == nil {
		panic(ErrNotInitialized)
	}

	if zyly == 3 {
		zyly = 2
	} else if zyly == 4 || zyly == 5 || zyly == 1 {
		zyly = 3
	}

	// find column to consider basing on zyly and gdzie
	columnIndex := 1
	for ; columnIndex < len(YDY_przekroj[0]); columnIndex++ {
		parts := strings.Split(YDY_przekroj[0][columnIndex], " ")
		liczbaZyl, _ := strconv.Atoi(parts[1])
		lokalizacja := parts[2]
		fmt.Println("lokalizacja:", lokalizacja)
		if zyly == liczbaZyl && gdzie == CableLocation(lokalizacja) {
			fmt.Printf("wybrałem kolumnę z liczbą żył %d i lokalizacją %s\n", zyly, gdzie)
			break
		}
	}

	// start from 1 to skip first row
	for i := 1; i < len(YDY_przekroj); i++ {
		thisRow := YDY_przekroj[i]
		fmt.Println("thisRow:", thisRow)
		thisNatezenie, _ := strconv.ParseFloat(thisRow[columnIndex], 64)
		nextCrossection, _ := strconv.ParseFloat(thisRow[0], 64)
		fmt.Println("i:", i, "thisNatezenie:", thisNatezenie, "nextCrossection:", nextCrossection)
		if thisNatezenie > I_ost {
			return nextCrossection, nil
		} else {
			fmt.Println("thisNatezenie <= I_ost", thisNatezenie, "<=", I_ost)
		}
	}

	return 0, ErrNotFound
}

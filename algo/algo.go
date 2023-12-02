package algo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ErrNotInitialized = fmt.Errorf("algo package not initialized")
	ErrNotFound       = fmt.Errorf("not found")
)

var (
	YDY_tolerance [][]string
	YDY_przekroj  [][]string
)

// Init loads data needed for the algorithm to work.
func Init() error {
	{
		file, err := os.Open("algo/tolerance_data.csv")
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
		file, err := os.Open("algo/YDY_przekroj.csv")
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
func CalcTemp(I_obl, temp float64) (float64, error) {
	if YDY_tolerance == nil {
		panic(ErrNotInitialized)
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
	A1 CableLocation = "A2"
	// przewód wielożyłowy w rurze instalacyjnej w ścianie izolowanej cieplnie
	A2 CableLocation = "A3"
	// przewód  w rurze instalacyjnej/listwie naściennej na ścianie murowanej
	// lub drewnianej
	B2 CableLocation = "B2"
	// przewód wielożyłowy w powietrzu (korytko perforowane, siatkowe)
	E CableLocation = "E2"
)

// MatchCrossection zwraca przekrój przewodu dla podanego I_ost, zyly i gdzie.
// I_ost = ostateczny prąd obciążenia
func MatchCrossection(I_ost float64, zyly int, gdzie CableLocation) {
	if YDY_przekroj == nil {
		panic(ErrNotInitialized)
	}

	// find column to consider basing on zyly and gdzie
	columnIndex := 1
	for ; columnIndex < len(YDY_przekroj[0]); columnIndex++ {
		parts := strings.Split(YDY_przekroj[0][columnIndex], " ")
		liczbaZyl, _ := strconv.Atoi(parts[1])
		lokalizacja := parts[2]

		if zyly == liczbaZyl && gdzie == CableLocation(lokalizacja) {
			fmt.Printf("wybrałem kolumnę z liczbą żył %d i lokalizacją %s\n", zyly, gdzie)
			break
		}
	}

	// start from 1 to skip first row
	for i := 1; i < len(YDY_przekroj[columnIndex]); i++ {
		thisRow := YDY_przekroj[i]
		fmt.Println("thisRow:", thisRow)
		thisNatezenie, _ := strconv.ParseFloat(thisRow[columnIndex], 64)
		nextCrossection, _ := strconv.ParseFloat(thisRow[0], 64)
		if thisNatezenie > I_ost {
			fmt.Printf("przekrój to %f mm^2\n", nextCrossection)
			return
		}
	}
}

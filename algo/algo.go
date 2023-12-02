package algo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var (
	ErrNotInitialized = fmt.Errorf("algo package not initialized")
	ErrNotFound       = fmt.Errorf("not found")
)

var YDYtolerance [][]string

// Init loads data needed for the algorithm to work.
func Init() error {
	file, err := os.Open("tolerance_data.csv")
	if err != nil {
		return fmt.Errorf("failed to load tolerance csv file: %w", err)
	}
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)
	YDYtolerance, err = csvReader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read from file into csv: %w", err)
	}

	return nil
}

func Calc(I_obl, tolerance float64) (I_ost float64) {
	I_ost = I_obl / (tolerance * 0.85)
	return
}

// CalcTemp returns I_ost for given I_obl and temp.
func CalcTemp(I_obl, temp float64) (float64, error) {
	if YDYtolerance == nil {
		panic(ErrNotInitialized)
	}

	// start from 1 to skip first column
	for i := 1; i < len(YDYtolerance); i++ {
		thisRow := YDYtolerance[i]
		thisTemp, _ := strconv.ParseFloat(thisRow[0], 64)
		nextTolerance, _ := strconv.ParseFloat(thisRow[1], 64)
		if thisTemp > temp {
			return Calc(I_obl, nextTolerance), nil
		}
	}

	return 0, ErrNotFound
}

func MatchCrossection

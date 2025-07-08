package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// statsFunc defines a generic statistical function
type statsFunc func(data []float64) float64

func sum(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func min(data []float64) float64 {
	min := data[0]
	for _, value := range data {
		if value < min {
			min = value
		}
	}
	return min
}

func max(data []float64) float64 {
	max := data[0]
	for _, value := range data {
		if value > max {
			max = value
		}
	}
	return max
}

func csv2float(r io.Reader, column int) ([]float64, error) {
	// Create the CSV reader used to read in data from a file
	cr := csv.NewReader(r)
	cr.ReuseRecord = true

	// Adjusting for 0 based index
	column--

	var data []float64

	for i := 0; ; i++ {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("cannot read data from file: %w", err)
		}
		if i == 0 {
			continue
		}

		// Skip the column if it's out of range
		if len(row) <= column {
			return nil, fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(row))
		}

		/// Parse the column value as a float
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}
		data = append(data, v)
	}

	return data, nil
}

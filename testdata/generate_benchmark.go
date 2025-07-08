package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
)

const (
	numFiles    = 1000
	numRows     = 2501 // includes the header row
	rowsPerFile = numRows - 1
)

func main() {
	benchmarkDir := filepath.Join("testdata", "benchmark")
	if err := os.MkdirAll(benchmarkDir, 0755); err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		return
	}

	for i := 0; i < numFiles; i++ {
		fileName := filepath.Join(benchmarkDir, fmt.Sprintf("file_%04d.csv", i))
		f, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", fileName, err)
			continue
		}

		w := csv.NewWriter(f)
		// Write header
		if err := w.Write([]string{"Col1", "Col2"}); err != nil {
			fmt.Printf("Failed to write header to %s: %v\n", fileName, err)
			f.Close()
			continue
		}
		// Write rows
		for j := 0; j < rowsPerFile; j++ {
			col1 := fmt.Sprintf("Data%d", j)
			col2 := strconv.Itoa(rand.Intn(99999))
			if err := w.Write([]string{col1, col2}); err != nil {
				fmt.Printf("Failed to write row to %s: %v\n", fileName, err)
				break
			}
		}
		w.Flush()
		if err := w.Error(); err != nil {
			fmt.Printf("Flush error for %s: %v\n", fileName, err)
		}
		f.Close()
		if (i+1)%100 == 0 {
			fmt.Printf("Generated %d files\n", i+1)
		}
	}
	fmt.Println("Benchmark CSV generation completed.")
}

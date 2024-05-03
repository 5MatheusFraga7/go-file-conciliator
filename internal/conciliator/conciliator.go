package conciliator

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Conciliator struct {
	InternalFile *os.File
	ExternalFile *os.File
}

func (c *Conciliator) Conciliate() {
	records := c.GetDataFromInternalFile()
	records2 := c.GetDataFromExternalFile()

	dataDiff := getBynaryFileDiff(records, records2)

	writeCSV("output_diff.csv", dataDiff)
}

func (c *Conciliator) GetDataFromInternalFile() [][]string {
	file := c.InternalFile
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	return records
}

func (c *Conciliator) GetDataFromExternalFile() [][]string {
	file := c.ExternalFile
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	return records
}

func getBynaryFileDiff(internalData [][]string, externalData [][]string) [][]string {
	var wg sync.WaitGroup
	var mu sync.Mutex
	diff := [][]string{}

	maxGoroutines := 100
	goroutines := make(chan struct{}, maxGoroutines)

	for i := 0; i < len(internalData); i++ {
		wg.Add(1)
		goroutines <- struct{}{}

		go func(idx int) {
			defer wg.Done()
			defer func() { <-goroutines }()

			if !binarySearch(internalData[idx], externalData) {
				mu.Lock()
				diff = append(diff, internalData[idx])
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()
	return diff
}

func binarySearch(value []string, array [][]string) bool {
	left := 0
	right := len(array) - 1

	target, _ := strconv.Atoi(value[0])

	for left <= right {
		mid := left + (right-left)/2
		num, _ := strconv.Atoi(array[mid][0])

		if num == target {
			return true
		} else if target < num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

func writeCSV(filename string, data [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

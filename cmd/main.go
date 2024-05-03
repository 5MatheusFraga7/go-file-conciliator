package main

import (
	"database-conciliator/internal/conciliator"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

func main() {

	// internalData, externalData := getDataToConciliate()

	// dataDiff := getBynaryFileDiff(internalData, externalData)

	// writeCSV("output_diff.csv", dataDiff)

	c := conciliator.Conciliator{File: getFileToConciliate()}
	c.Conciliate()

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

func getFileToConciliate() *os.File {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envFile := filepath.Join(dir, ".env")

	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	err = godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	filePath := os.Getenv("FILE_PATH")

	file, err := os.Open(filePath)

	return file
}

func getDataToConciliate() ([][]string, [][]string) {
	internalData := [][]string{
		{"1"},
		{"2"},
		{"3"},
		{"4"},
		{"5"},
		{"6"},
		{"7"},
		{"8"},
		{"19"},
		{"20"},
	}
	externalData := [][]string{
		{"1"},
		{"2"},
		{"3"},
		{"4"},
		{"5"},
		{"6"},
		{"7"},
		{"8"},
		{"9"},
		{"10"},
	}

	return internalData, externalData
}

package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func main() {
	internalData := getFileData("data1.csv")
	externalData := getFileData("data2.csv")

	dataDiff := getBynaryFileDiff(internalData, externalData)

	writeCSV("output_diff.csv", dataDiff)
}

func getBynaryFileDiff(internalData [][]string, externalData [][]string) [][]string {

	diff := [][]string{}

	for i := 0; i < len(internalData); i++ {
		if !binarySearch(internalData, externalData[i][0]) {
			diff = append(diff, internalData[i])
		}
	}
	return diff
}

func binarySearch(array [][]string, value string) bool {

	left := 0
	right := len(array) - 1

	for left <= right {
		mid := left + (right-left)/2
		num, _ := strconv.Atoi(array[mid][0])
		target, _ := strconv.Atoi(value)

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

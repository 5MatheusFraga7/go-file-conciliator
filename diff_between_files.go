package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
)

func diff() {
	internalData := getFileData("data1.csv")
	externalData := getFileData("data2.csv")

	combinedData := joinData(internalData, externalData)
	sortedData := sortData(combinedData)

	conciliateData(sortedData)
}

func getFileData(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	return records
}

func joinData(internalData [][]string, externalData [][]string) [][]string {
	combinedData := append(internalData[:], externalData[:]...)
	return combinedData
}

func sortData(array [][]string) [][]string {
	sort.Slice(array, func(i, j int) bool {
		return array[i][0] < array[j][0]
	})
	return array
}

func conciliateData(array [][]string) {
	result := []string{}
	result = concilia(array, 0, result)
	fmt.Println(result)
}

func concilia(array [][]string, i int, result []string) []string {
	if i+1 >= len(array) {
		return result
	}

	if array[i][0] == array[i+1][0] {
		return concilia(array, i+2, result)
	} else {
		result = append(result, array[i+1][0])
		return concilia(array, i+1, result)
	}
}

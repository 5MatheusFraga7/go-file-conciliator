package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	internalData := getFileData("data1.csv")
	externalData := getFileData("data2.csv")

	err := writeCSV("diff.csv", diffFiles(internalData, externalData))
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo CSV:", err)
		return
	}
	fmt.Println("Dados foram escritos com sucesso no arquivo diff.csv")

}

func diffFiles(baseA [][]string, baseB [][]string) [][]string {
	presenceB := make(map[string]bool)
	for _, row := range baseB {
		for _, val := range row {
			presenceB[val] = true
		}
	}

	diff := [][]string{}
	for _, row := range baseA {
		isDifferent := false
		for _, val := range row {
			if !presenceB[val] {
				isDifferent = true
				break
			}
		}
		if isDifferent {
			diff = append(diff, row)
		}
	}

	return diff
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

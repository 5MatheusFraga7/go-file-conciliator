package main

import (
	"database-conciliator/internal/conciliator"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {

	internalFile, externalFile := getFilesToConciliate()

	c := conciliator.Conciliator{InternalFile: internalFile, ExternalFile: externalFile}
	c.Conciliate()
}

func getFilesToConciliate() (*os.File, *os.File) {
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

	filePath := os.Getenv("INTERNAL_FILE_PATH")
	filePath2 := os.Getenv("EXTERNAL_FILE_PATH")

	file1, err := os.Open(filePath)
	file2, err := os.Open(filePath2)

	return file1, file2
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

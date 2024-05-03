package conciliator

import (
	"database-conciliator/internal/db"
	"database-conciliator/internal/db/adapters"
	"encoding/csv"
	"fmt"
	"os"
)

type Conciliator struct {
	File *os.File
}

func (c *Conciliator) Conciliate() {
	startDatabase()

	records := c.GetDataFromInternalFile()

	for _, row := range records {
		fmt.Println(row)
	}

}

func (c *Conciliator) GetDataFromInternalFile() [][]string {
	file := c.File
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	return records
}

func startDatabase() {
	postgresAdapter := adapters.NewPostgreSQLAdapter()
	err := db.OpenConnectionToDatabase(postgresAdapter)
	defer db.CloseConnectionToDatabase(postgresAdapter)

	if err != nil {
		fmt.Println(err)
	}

}

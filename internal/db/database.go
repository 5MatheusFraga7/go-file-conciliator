package db

import "fmt"

type Database interface {
	Connect() error
	Close() error
}

func OpenConnectionToDatabase(database Database) error {
	err := database.Connect()
	if err != nil {
		fmt.Println("EROOOOOOO!!!!!: %v", err)
		return err
	}

	fmt.Println("Conexão aberta!!!!!: %v", database)

	return nil
}

func CloseConnectionToDatabase(database Database) {
	database.Close()
}

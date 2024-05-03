package adapters

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func GetConnectionString() string {
	config := GetDatabaseConfigEnvs()

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.Database)
}

func GetDatabaseConfigEnvs() Config {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envFile := filepath.Join(dir, ".env")

	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	// Carrega o arquivo .env
	err = godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	return Config{Host: host, Port: port, Username: username, Password: password, Database: database}

}

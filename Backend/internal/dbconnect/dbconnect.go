package dbconnect

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() (*sql.DB, error) {
	// Connection string
	psqlInfo, err := getConnectionParamInfo()

	if err != nil {
		return nil, fmt.Errorf("error with connection parameters: %w", err)
	}
	// Open a connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func getConnectionParamInfo() (string, error) {
	file, err := os.ReadFile(".env/SQLConnectionParameters.json")
	if err != nil {
		return "", fmt.Errorf("error reading config file: %w", err)
	}

	var config DBConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		return "", fmt.Errorf("error parsing config file: %w", err)
	}

	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	return psqlInfo, nil
}

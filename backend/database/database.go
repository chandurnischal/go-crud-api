package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

var DB *sql.DB

func Connect() error {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	for i := 0; i < 10; i++ {
		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			return fmt.Errorf("failed to connect to database (attempt %d): %v", i+1, err)
		} else if err = DB.Ping(); err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return fmt.Errorf("failed to connect to database after multiple attempts: %v", err)
	}
	return nil
}

func UserExists(name string) bool {
	user, err := RetrieveByName(name)

	if err != nil {
		return false
	}

	if user.Name != "" {
		return false
	}

	return true

}

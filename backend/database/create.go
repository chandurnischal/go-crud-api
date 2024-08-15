package database

import "fmt"

func CreateUser(name, location string) error {

	if UserExists(name) {
		return fmt.Errorf("user exists in database")
	}

	_, err := DB.Exec("INSERT INTO users (name, location) VALUES (?, ?)", name, location)

	return err
}

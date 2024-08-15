package database

import (
	"fmt"
	"strings"
)

func RetrieveByName(name string) (User, error) {
	var user User
	query := fmt.Sprintf("SELECT * FROM users WHERE LOWER(name) LIKE '%%%s%%'", strings.ToLower(name))

	row := DB.QueryRow(query)

	err := row.Scan(&user.ID, &user.Name, &user.Location)

	if err != nil {
		return user, err
	}

	err = row.Err()

	if err != nil {
		return user, err
	}

	return user, err
}

func RetrieveAll() ([]User, error) {
	var users []User

	query := "SELECT * FROM users"

	rows, err := DB.Query(query)

	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Location)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return users, err
	}

	return users, err
}

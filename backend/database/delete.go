package database

import (
	"strings"
)

func DeleteUser(name string) error {

	_, err := DB.Exec("DELETE FROM users WHERE LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%")
	if err != nil {
		return err
	}

	return nil
}

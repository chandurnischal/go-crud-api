package database

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

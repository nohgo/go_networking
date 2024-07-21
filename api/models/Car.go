package models

// Represents a car in the database
type Car struct {
	Id    int    `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

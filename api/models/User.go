package models

// Represents a user in the database
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Cars     []Car  `json:"cars"`
}

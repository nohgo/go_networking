package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Cars     []Car  `json:"cars"`
}

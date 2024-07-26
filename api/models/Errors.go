package models

type UserError struct {
}

func (e *UserError) Error() string {
	return "Invalid input"
}

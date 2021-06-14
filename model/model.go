package model

//Users ...
type Users struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Password string `json:"phone" db:"passsword" `
	Email    string `json:"email" db:"email"`
}

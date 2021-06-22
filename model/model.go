package model

//Users ...
type Users struct {
	ID       int    `json:"ID" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name     string `json:"Name,omitempty" db:"Name"`
	Password string `json:"Password,omitempty" db:"Passsword" `
	Email    string `gorm:"unique" json:"Email,omitempty" db:"Email"`
}

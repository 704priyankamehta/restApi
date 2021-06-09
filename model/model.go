package model

//Users ...
type Users struct {
	ID    int    `gorm:"primaryKey" ; autoIncrement:true`
	Name  string `json:"name" db:"name"`
	Phone int    `json:"phone" db:"phone" `
}

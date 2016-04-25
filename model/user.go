package model

// User holds user personal data
type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstname" sql:"not null"`
	LastName  string `json:"lastname" sql:"not null"`
	Email     string `json:"email" sql:"not null;unique"`
	Password  string `json:"-" sql:"not null"`
}

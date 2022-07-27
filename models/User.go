package models

type User struct {
	ID        int    `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Name      string `json:"name"`
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
	userName  string `json:"userName"`
	Email     string `json:"email"`
}

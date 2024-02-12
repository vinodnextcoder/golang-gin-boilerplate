package models

type User struct {
	Id       uint   `json:"Id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Id       uint   `json:"Id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

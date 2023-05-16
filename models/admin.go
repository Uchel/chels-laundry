package models

type Admin struct {
	ID       string `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Photo    string `json:"photo"`
	Password string `json:"password" form:"password"`
}

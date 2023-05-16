package models

type Customer struct {
	// ID       string `json:"" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	PassCode string `json:"pass_code" validate:"required"`
}

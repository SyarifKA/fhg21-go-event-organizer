package models

type Users struct {
	Id       int     `json:"id"`
	Email    string  `json:"email"`
	Password string  `json:"-"`
	Username *string `json:"username"`
}

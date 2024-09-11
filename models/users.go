package models

type Users struct {
	Id       int     `json:"id"`
	Email    string  `json:"email"`
	Password string  `json:"-"`
	Username *string `json:"username"`
}

type UserUpdate struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	// Password string  `json:"-"`
	Username *string `json:"username"`
}

type UpdatePassword struct {
	OldPassword     string
	NewPassword     string
	ConfirmPassword string
}

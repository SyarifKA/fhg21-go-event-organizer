package dtos

type User struct {
	Id       int     `json:"id"`
	Email    string  `json:"email" form:"email" binding:"required,email"`
	Password string  `json:"password" form:"password" binding:"required,min=8"`
	Username *string `json:"username" binding:"omitempty"`
}

type Password struct {
	// OldPassword     string `json:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" form:"newPassword" db:"password"`
	// ConfirmPassword string `json:"confirmPassword" form:"confirmPassword"`
}

type FormUser struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
	RoleId   int    `form:"roleId"`
}

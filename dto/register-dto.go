package dto

//Register struct for validate payload from register user data
type RegisterDTO struct {
	FullName        string `json:"full_name" form:"name" binding:"required,min=2"`
	Username        string `json:"username" form:"username" binding:"required,email" `
	Password        string `json:"password" form:"password" binding:"required,min=12"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,min=12"`
}

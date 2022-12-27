package dto

//LoginDTO struct for validate payload from login url
type LoginDTO struct {
	Username string `json:"username" form:"username" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

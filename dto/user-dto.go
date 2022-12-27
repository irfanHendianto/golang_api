package dto

//UserUpdateDTO struct for validate payload from update user data
type UserUpdateDTO struct {
	ID              uint64 `json:"id" form:"id"`
	FullName        string `json:"full_name" form:"full_name" binding:"required,min=2"`
	Username        string `json:"username" form:"username" binding:"required,email"`
	Password        string `json:"password,omitempty" form:"password,omitempty,min=12"`
	ConfirmPassword string `json:"ConfirmPassword,omitempty" form:"ConfirmPassword,omitempty,min=12"`
}

// DleteUSerDTO struct for validate payload from delete user
type DeleteUserDTO struct {
	ID int64 `json:"id" form:"id"`
}

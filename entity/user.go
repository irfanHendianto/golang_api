package entity

// Represent user table
type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	FullName string `gorm:"type:varchar(255)" json:"full_name"`
	Username string `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	Password string `gorm:"->;<-;not null " json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}

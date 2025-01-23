package model

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	IsActive  bool   `json:"is_active"`
	RoleID    string `json:"role_id"`
	Role      Role   `gorm:"foreignKey:RoleID" json:"role"`
}

type UserRegiterInput struct {
	Username string `json:"username" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

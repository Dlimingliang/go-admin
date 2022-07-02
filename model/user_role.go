package model

type UserRole struct {
	UserID int `gorm:"column:user_id"`
	RoleID int `gorm:"column:role_authority_id"`
}

func (s *UserRole) TableName() string {
	return "user_roles"
}

package entity

import "github.com/marcleonschulz/carSearchApi/pkg/models"

type User struct {
	Username  string     `gorm:"column:username;type:varchar(100)"`
	Password  string     `gorm:"column:password;type:varchar(200)"`
	Email     string     `gorm:"primaryKey;column:email;type:varchar(100)"`
	IsActive  bool       `gorm:"column:is_active;type:boolean"`
	UserRoles []UserRole `gorm:"ForeignKey:email;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// ToResponse converts User to UserResponse
func (user *User) ToResponse() models.UserResponse {
	var roles []string
	for _, userRole := range user.UserRoles {
		roles = append(roles, userRole.Role)
	}
	return models.UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Roles:    roles,
	}
}

func (User) TableName() string {
	return "tb_user"
}

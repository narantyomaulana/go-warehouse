package model

import "time"

type UserRole struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	RoleId    uint      `json:"role_id" gorm:"not null"`
	User      User      `gorm:"foreignKey:UserId"`
	Role      Role      `gorm:"foreignKey:RoleId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tabler interface {
	TableName() string
}

func (UserRole) TableName() string {
	return "user_role"
}

package Models

import (
	"time"
)

type UserModel struct {
	Id        uint      `gorm:"AUTO_INCREMENT;primaryKey;unique;not nul" json:"id"`
	UserName  string    `gorm:"unique;not nul" json:"username"`
	Email     string    `gorm:"unique;not nul" json:"email"`
	Password  string    `gorm:"not nul" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *UserModel) TableName() string {
	return "user"
}

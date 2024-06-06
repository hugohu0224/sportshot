package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:char(36);primaryKey"`
	Username     string    `gorm:"unique;not null"`
	Password     string    `gorm:"not null"`
	RefreshToken string
	Active       bool `gorm:"default:true"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

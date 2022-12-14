package models

import (
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/irfan44/task-5-vix-btpns-IrfanNurghiffariM/helpers/hashpass"
)

type User struct {
	ID        string    `gorm:"primary_key; unique" json:"id"`
	Username  string    `gorm:"size:255;not null;" json:"username"`
	Email     string    `gorm:"size:100;not null; unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Photos    Photo     `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"photos"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) Initialize() {
	u.ID = uuid.New().String()
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) HashPassword() {
	hashedPassword, err := hashpass.HashPassword(u.Password)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)
}

func (u *User) CheckPasswordHash(password string) bool {
	return hashpass.CheckPasswordHash(password, u.Password)
}

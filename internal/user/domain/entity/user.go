package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:255;not null"`
	Email    string `gorm:"uniqueIndex;size:255;not null"`
	Password string `gorm:"size:255;not null"`
}

func (u *User) TableName() string {
	return "users"
}

func NewUser(username, email, password string) (*User, error) {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
	}, nil
}

func (u *User) UpdateEmail(newEmail string) error {
	u.Email = newEmail
	return nil
}

func (u *User) UpdatePassword(newPassword string) error {
	u.Password = newPassword
	return nil
}

func (u *User) ValidatePassword(password string) bool {
	return u.Password == password
}

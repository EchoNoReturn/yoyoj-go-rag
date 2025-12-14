package entity

import (
	"net/mail"
	"regexp"
	user_errors "yoyoj-go-rag/internal/user/domain/errors"
	"yoyoj-go-rag/internal/user/domain/valueobject"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:255;not null"`
	Email    string `gorm:"uniqueIndex;size:255;not null"`
	PasswordHash string `gorm:"size:255;not null"`
}

func (u *User) TableName() string {
	return "users"
}

func ValidUsernameType(username string) (bool, error) {
	if username == "" {
		return false, user_errors.ErrInvalidUsername
	}
	return true, nil
}

func ValidEmailType(email string) (bool, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, user_errors.ErrInvalidEmail.WithDetail("email", email)
	}
	return err == nil, err
}

func ValidPasswordType(password string) (bool, error) {
	// 密码验证配置（可根据需求调整）
	const (
		minLength = 8  // 密码最小长度
		maxLength = 72 // bcrypt 限制最大 72 字符
	)
	if len(password) < minLength || len(password) > maxLength {
		return false, user_errors.ErrInvalidPasswordLength
	}
	// 判断是否为大写字母、小写字母、数字的组合
	var passwordRegex = regexp.MustCompile(`^[A-Za-z0-9]+$`)
	if !passwordRegex.MatchString(password) {
		return false, user_errors.ErrInvalidPassword
	}
	return true, nil
}

func NewUser(username, email, password string) (*User, error) {
	// 1. 验证用户名
	if valid, err := ValidUsernameType(username); !valid {
		return nil, err
	}

	// 2. 验证邮箱
	if valid, err := ValidEmailType(email); !valid {
		return nil, err
	}

	// 3. 验证密码
	if valid, err := ValidPasswordType(password); !valid {
		return nil, err
	}

	// 4. ✅ 加密密码
	hashedPassword, err := valueobject.NewHashedPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		Username: username,
		Email:    email,
		PasswordHash: hashedPassword.String(),
	}, nil
}

func (u *User) UpdateEmail(newEmail string) error {
	if valid, err := ValidEmailType(newEmail); !valid {
		return err
	}
	u.Email = newEmail
	return nil
}

func (u *User) UpdatePassword(newPassword string) error {
	// 1. 验证新密码格式
	if valid, err := ValidPasswordType(newPassword); !valid {
		return err
	}
	// 2. 加密新密码
	hashedPassword, err := valueobject.NewHashedPassword(newPassword)
	if err != nil {
		return err
	}
	u.PasswordHash = hashedPassword.String()
	return nil
}

func (u *User) ValidatePassword(plainPassword string) bool {
	// ✅ 使用 bcrypt 验证
	hashedPassword := valueobject.FromHash(u.PasswordHash)
	return hashedPassword.Verify(plainPassword)
}

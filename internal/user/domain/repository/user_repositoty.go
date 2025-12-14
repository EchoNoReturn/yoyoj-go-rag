package repository

import "yoyoj-go-rag/internal/user/domain/entity"

type UserRepository interface {
	Save(user *entity.User) error
	FindByID(id uint) (*entity.User, error)
	GetRegisterSchema() string
	Delete(id uint) error
}
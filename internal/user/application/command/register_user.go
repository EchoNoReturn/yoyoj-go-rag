package command

import "yoyoj-go-rag/internal/user/domain/repository"

type UpdateUserCommand struct {
	UserID      uint
	NewEmail    string
	NewPassword string
}

type UpdateUserHandler struct {
	userRepo repository.UserRepository
}
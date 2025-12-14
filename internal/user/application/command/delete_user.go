package command

import "yoyoj-go-rag/internal/user/domain/repository"

type DeleteUserCommand struct {
	UserID uint
}

type DeleteUserHandler struct {
	userRepo repository.UserRepository
}

func NewDeleteUserHandler(userRepo repository.UserRepository) *DeleteUserHandler {
	return &DeleteUserHandler{
		userRepo: userRepo,
	}
}

// Additional methods for DeleteUserHandler can be added here
func (h *DeleteUserHandler) Delete(cmd *DeleteUserCommand) error {
	// TODO 实现删除用户逻辑
	return nil
}
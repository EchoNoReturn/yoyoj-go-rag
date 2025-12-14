package command

import (
	"yoyoj-go-rag/internal/user/domain/entity"
	"yoyoj-go-rag/internal/user/domain/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type RegisterUserCommand struct {
	Username string
	Email    string
	Password string
}

type RegisterUserHandler struct {
	userRepo repository.UserRepository
}

func NewRegisterUserHandler(userRepo repository.UserRepository) *RegisterUserHandler {
	return &RegisterUserHandler{
		userRepo: userRepo,
	}
}

// Additional methods for RegisterUserHandler can be added here

// 注册用户。
// 所有的用户都通过这个方法实现注册，包括普通用户和管理员用户。管理员是被授予管理权限的用户。
func (h *RegisterUserHandler) Register(cmd *RegisterUserCommand) error {
	user, err := entity.NewUser(cmd.Username, cmd.Email, cmd.Password)
	if err != nil {
		return err
	}
	log.Infof("合法用户: %s", user.Username)
	if err = h.userRepo.Save(user); err != nil {
		return err
	}
	return nil
}

func (h *RegisterUserHandler) Handler(cmd *RegisterUserCommand) error {
	schema := h.userRepo.GetRegisterSchema()

	switch schema {
	case "public":
		return h.Register(cmd)
	case "review":
		// TODO 完善审核注册规则
		return nil;
	case "invite":
		// TODO 完善邀请注册规则
		return nil;
	case "closed":
		return fiber.ErrForbidden
	default:
		return nil
	}
}

package infrastructure

import (
	"yoyoj-go-rag/internal/user/domain/entity"
	"yoyoj-go-rag/internal/user/domain/repository"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(user *entity.User) error {
	// TODO 实现数据保存操作
	return nil
}

func (r *PostgresUserRepository) FindByID(id uint) (*entity.User, error) {
	user := entity.User{}
	// TODO 实现数据库查询工作
	return &user, nil
}

// 只有实例上有 DeleteAt 字段的模型才能实现软删除, gorm.Model 已经包含该字段
func (r *PostgresUserRepository) Delete(id uint) error {
	// TODO 实现删除用户逻辑
	user, err := r.FindByID(id)
	if err != nil {
		return err
	}
	return r.db.Delete(user).Error
}

func (r *PostgresUserRepository) DeleteForever(id uint) error {
	// TODO 实现永久删除用户逻辑
	user, err := r.FindByID(id)
	if err != nil {
		return err
	}
	return r.db.Unscoped().Delete(user).Error
}

func (r *PostgresUserRepository) GetRegisterSchema() string {
	// TODO 查询数据库常量表
	return "public"
}

var _ repository.UserRepository = (*PostgresUserRepository)(nil)

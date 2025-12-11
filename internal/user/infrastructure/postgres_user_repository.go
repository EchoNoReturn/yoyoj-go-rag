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
	return nil
}

func (r *PostgresUserRepository) FindByID(id uint) (*entity.User, error) {
	user := entity.User{}
	return &user, nil	
}

var _ repository.UserRepository = (*PostgresUserRepository)(nil)
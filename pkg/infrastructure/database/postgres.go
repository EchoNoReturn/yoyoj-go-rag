package database

import (
	"yoyoj-go-rag/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewPostgresConnection(cfg *configs.PostgresConfig) (*PostgresDB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresDB{DB: db}, nil
}

func (p *PostgresDB) Close() error {
	if sqlDB, err := p.DB.DB(); err != nil {
		return err
	} else {
		return sqlDB.Close()
	}
}

package storage

import (
	"gorm.io/gorm"

	"github.com/kavladst/products_api/internal/app/configuration"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(config *configuration.Configuration) (*Storage, error) {
	db, err := InitDB(config.DBHost, config.DBPort, config.DBName, config.DBUser, config.DBPassword)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

package storage

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	Id string `gorm:"primaryKey"`
}

func (s *Storage) IsUserExists(userId string) bool {
	err := s.db.First(&User{Id: userId}).Error
	if err == nil {
		return true
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	panic(err)
}

func (s *Storage) CreateUser(userId string) error {
	err := s.db.Create(&User{Id: userId}).Error
	if err == nil {
		return err
	}
	if err.Error() == "ERROR: duplicate key value violates unique constraint \"users_pkey\" (SQLSTATE 23505)" {
		return errors.New("user with this ID is already exists")
	}
	panic(err)
}

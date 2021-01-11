package storage

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id      uuid.UUID `gorm:"default:uuid_generate_v4()"`
	User    User      `gorm:"foreignKey:UserId"`
	UserId  string    `gorm:"not null;<-:create;uniqueIndex:products_user_id_offer_id_key"`
	OfferId string    `gorm:"not null;<-:create;uniqueIndex:products_user_id_offer_id_key"`
	Name    string    `gorm:"not null"`
	Price   int       `gorm:"not null"`
	Quality int       `gorm:"not null"`
}

func (s Storage) GetProductsByCondition(conditionString string) []Product {
	products := []Product{}
	err := s.db.Where(conditionString).Find(&products).Error
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return products
	}
	panic(err)
}

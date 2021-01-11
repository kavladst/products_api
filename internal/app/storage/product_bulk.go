package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductsForBulk struct {
	user                User
	availableProducts   []Product
	unavailableOfferIds []string
	offerIds            map[string]bool
}

func (s Storage) NewProductsForBulk(userId string) *ProductsForBulk {
	return &ProductsForBulk{user: User{Id: userId}, availableProducts: []Product{}, offerIds: map[string]bool{}}
}

func (p *ProductsForBulk) AppendAvailableProduct(offerId string, name string, price int, quality int) {
	if _, isInMap := p.offerIds[offerId]; isInMap {
		return
	}
	p.offerIds[offerId] = true
	newProduct := Product{User: p.user, OfferId: offerId, Name: name, Price: price, Quality: quality}
	p.availableProducts = append(p.availableProducts, newProduct)
}

func (p *ProductsForBulk) AppendUnavailableOfferId(offerId string) {
	p.unavailableOfferIds = append(p.unavailableOfferIds, offerId)
}

func (s *Storage) PushProductsBulk(products *ProductsForBulk) (int, int, int) {
	countNewRecords := 0

	tx := s.db.Begin()

	if len(products.availableProducts) > 0 {
		var countAllOldRecords int64
		panicIfError(
			tx.Table("products").Count(&countAllOldRecords),
		)
		panicIfError(
			tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}, {Name: "offer_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"name", "price", "quality"}),
			}).Create(&(products.availableProducts)),
		)
		var countAllNewRecords int64
		panicIfError(
			tx.Table("products").Count(&countAllNewRecords),
		)
		countNewRecords = int(countAllNewRecords - countAllOldRecords)
	}
	countDeletedRecords := 0
	for _, offerId := range products.unavailableOfferIds {
		res := tx.Where("user_id = ? AND offer_id = ?", products.user.Id, offerId).Delete(&Product{})
		panicIfError(res)
		countDeletedRecords += int(res.RowsAffected)
	}

	tx.Commit()
	return countNewRecords, len(products.availableProducts) - countNewRecords, countDeletedRecords
}

func panicIfError(tx *gorm.DB) {
	if tx.Error != nil {
		tx.Rollback()
		panic(tx.Error)
	}
}

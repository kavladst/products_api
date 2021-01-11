package api

import (
	"errors"
	"sort"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func (a *Api) creteUser(userId string) map[string]interface{} {
	err := a.storage.CreateUser(userId)
	if err != nil {
		return createErrorResponse("creating user", err)
	}
	return map[string]interface{}{
		"status": "ok",
		"id":     userId,
	}
}

func (a *Api) getProducts(userId string, offerId string, search string) []map[string]interface{} {
	conditionString := ""
	if userId != "" {
		conditionString += "user_id = " + userId
	}
	if offerId != "" {
		if len(conditionString) > 0 {
			conditionString += " AND "
		}
		conditionString += "offer_id = " + offerId
	}
	if search != "" {
		if len(conditionString) > 0 {
			conditionString += " AND "
		}
		conditionString += "name LIKE '%" + search + "%'"
	}
	products := a.storage.GetProductsByCondition(conditionString)
	data := make([]map[string]interface{}, len(products))
	for i, product := range products {
		data[i] = map[string]interface{}{
			"offer_id": product.OfferId,
			"name":     product.Name,
			"price":    product.Price,
			"quality":  product.Quality,
		}
	}
	return data
}

func (a *Api) updateProductsByXLSX(xlsx *excelize.File, userId string) map[string]interface{} {
	if !a.storage.IsUserExists(userId) {
		return createErrorResponse("user with input ID does not exist", nil)
	}
	productsForBulk := a.storage.NewProductsForBulk(userId)
	notProcessedIndexes := map[string][]int{}
	rows := xlsx.GetRows("Sheet1")
	for rowIndex, row := range rows {
		if len(row) < 5 {
			errMessage := "5 fields are required in row"
			notProcessedIndexes[errMessage] = append(notProcessedIndexes[errMessage], rowIndex)
			continue
		}
		offerId, name, price, quality, available, err := convertProductRow(row[0:5])
		if err != nil {
			notProcessedIndexes[err.Error()] = append(notProcessedIndexes[err.Error()], rowIndex)
		} else if available {
			productsForBulk.AppendAvailableProduct(offerId, name, price, quality)
		} else {
			productsForBulk.AppendUnavailableOfferId(offerId)
		}
	}

	countCreatedRecords, countUpdatedRecords, countDeletedRecords := a.storage.PushProductsBulk(productsForBulk)

	errorNames := make([]string, 0, len(notProcessedIndexes))
	for errorName := range notProcessedIndexes {
		errorNames = append(errorNames, errorName)
	}
	sort.Strings(errorNames)
	notProcessedInfo := make([]map[string]interface{}, 0, len(notProcessedIndexes))
	for _, errorName := range errorNames {
		notProcessedInfo = append(notProcessedInfo, map[string]interface{}{
			"error": errorName,
			"indexes_in_xlsx": notProcessedIndexes[errorName],
		})
	}
	info := map[string]interface{}{
		"count_new_products":     countCreatedRecords,
		"count_updated_products": countUpdatedRecords,
		"count_deleted_products": countDeletedRecords,
		"not_processed":          notProcessedInfo,
	}
	return info
}

func convertProductRow(row []string) (offerId string, name string, price int, quality int, available bool, err error) {
	if len(row) != 5 {
		err = errors.New("5 fields are required in row")
		return
	}
	offerId = row[0]
	if offerId == "" {
		err = errors.New("offer ID is required")
		return
	}
	name = row[1]
	if name == "" {
		err = errors.New("name is required")
		return
	}
	price, err = strconv.Atoi(row[2])
	if err != nil {
		err = errors.New("price is not converted to int")
		return
	}
	if price < 0 {
		err = errors.New("price is less than 0")
		return
	}
	quality, err = strconv.Atoi(row[3])
	if err != nil {
		err = errors.New("quality is not converted to int")
		return
	}
	if quality < 0 {
		err = errors.New("quality is less than 0")
		return
	}
	available, err = strconv.ParseBool(row[4])
	if err != nil {
		err = errors.New("available is not converted to bool")
		return
	}
	return
}

func createErrorResponse(message string, inputError error) map[string]interface{} {
	messageError := message
	if inputError != nil {
		messageError += ". " + inputError.Error()
	}
	return map[string]interface{}{"error": messageError}
}

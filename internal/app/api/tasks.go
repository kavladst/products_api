package api

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/google/uuid"
)

func (a *Api) taskCreateUser(userId string) uuid.UUID {
	taskId := a.taskManager.CreateTask()
	go func() {
		defer func(res map[string]interface{}) {
			a.taskManager.PutTaskResult(taskId, res)
		}(a.creteUser(userId))
	}()
	return taskId
}

func (a *Api) taskGetProducts(userId string, offerId string, search string) uuid.UUID {
	taskId := a.taskManager.CreateTask()
	go func() {
		defer func(res []map[string]interface{}) {
			a.taskManager.PutTaskResult(taskId, map[string]interface{}{"data": res, "count": len(res)})
		}(a.getProducts(userId, offerId, search))
	}()
	return taskId
}

func (a *Api) taskUpdateProducts(xlsx *excelize.File, userId string) uuid.UUID {
	taskId := a.taskManager.CreateTask()
	go func() {
		defer func(res interface{}) {
			a.taskManager.PutTaskResult(taskId, res)
		}(a.updateProductsByXLSX(xlsx, userId))
	}()
	return taskId
}

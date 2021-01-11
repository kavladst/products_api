package api

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *Api) initRouter() {
	router := gin.Default()
	v1 := router.Group("/v1")
	users := v1.Group("/users")
	products := v1.Group("/products")
	tasks := v1.Group("/tasks")

	users.POST("/", a.createUser)

	products.GET("/", a.getAllProducts)
	products.POST("/update/xlsx", a.updateProducts)

	tasks.GET("/", a.getTaskResult)

	a.router = router
}

func (a *Api) createUser(ctx *gin.Context) {
	userId := ctx.PostForm("id")
	if userId == "" {
		ctx.JSON(400, map[string]string{"error": "id is required"})
		return
	}
	taskId := a.taskCreateUser(userId)
	ctx.JSON(200, map[string]uuid.UUID{"task_id": taskId})
}

func (a *Api) getAllProducts(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	offerId := ctx.Query("offer_id")
	search := ctx.Query("search")
	taskId := a.taskGetProducts(userId, offerId, search)
	ctx.JSON(200, map[string]uuid.UUID{"task_id": taskId})
}

func (a *Api) updateProducts(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if userId == "" {
		ctx.JSON(400, map[string]string{"error": "user_id is required"})
		return
	}
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(400, map[string]string{"error": "failed to read file"})
		return
	}
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		ctx.JSON(400, map[string]string{"error": "file is not converted to xlsx"})
		return
	}
	taskId := a.taskUpdateProducts(xlsx, userId)
	ctx.JSON(200, map[string]uuid.UUID{"task_id": taskId})
}

func (a *Api) getTaskResult(ctx *gin.Context) {
	taskId, err := uuid.Parse(ctx.Query("task_id"))
	if err != nil {
		ctx.JSON(400, map[string]interface{}{"error": "task_id is not convert to UUID"})
		return
	}
	ctx.JSON(200, map[string]interface{}{
		"result": a.taskManager.GetTaskResult(taskId),
	})
}

package api

import (
	"github.com/gin-gonic/gin"

	"github.com/kavladst/products_api/internal/app/configuration"
	"github.com/kavladst/products_api/internal/app/storage"
	"github.com/kavladst/products_api/internal/app/task_manager"
)

type Api struct {
	router      *gin.Engine
	storage     *storage.Storage
	Config      *configuration.Configuration
	taskManager *task_manager.TaskManager
}

func NewApi() (*Api, error) {
	apiConfig, err := configuration.NewConfig()
	if err != nil {
		return nil, err
	}

	apiStorage, err := storage.NewStorage(apiConfig)
	if err != nil {
		return nil, err
	}

	newApi := &Api{storage: apiStorage, Config: apiConfig, taskManager: task_manager.NewTaskManager()}
	newApi.initRouter()

	return newApi, nil
}

func (a *Api) Run() error {
	return a.router.Run(a.Config.AppHost + ":" + a.Config.AppPort)
}

package task_manager

import (
	"sync"

	"github.com/google/uuid"
)

type TaskManager struct {
	Tasks map[uuid.UUID]interface{}
	Mutex *sync.Mutex
}

func NewTaskManager() *TaskManager {
	taskManager := TaskManager{}
	taskManager.Tasks = map[uuid.UUID]interface{}{}
	taskManager.Mutex = &sync.Mutex{}
	return &taskManager
}

func (t *TaskManager) CreateTask() uuid.UUID {
	taskId := uuid.New()
	t.Tasks[taskId] = nil
	return taskId
}

func (t *TaskManager) GetTaskResult(taskId uuid.UUID) interface{} {
	return t.Tasks[taskId]
}

func (t *TaskManager) PutTaskResult(taskId uuid.UUID, result interface{}) {
	t.Mutex.Lock()
	t.Tasks[taskId] = result
	t.Mutex.Unlock()
}

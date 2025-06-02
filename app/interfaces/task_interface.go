package interfaces

import (
	"github.com/JoyceEllenNeryTeles/test/tasks-api/app/domain/entity"
	"github.com/JoyceEllenNeryTeles/test/tasks-api/app/usecase"
	"github.com/gin-gonic/gin"
)

var _ TaskUsecaseInterface = (*usecase.TaskUsecase)(nil)

type TaskUsecaseInterface interface {
	GetTasks() ([]entity.Task, error)
	AddTask(ctx *gin.Context, task entity.Task) (int64, error)
	FindTaskById(id int64) (*entity.Task, error)
	FindTaskByOwner(owner string) ([]entity.Task, error)
	UpdateTask(ctx *gin.Context, task entity.Task) error
	DeleteTask(id int64) error
}

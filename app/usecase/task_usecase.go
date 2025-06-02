package usecase

import (
	"github.com/JoyceEllenNeryTeles/test/tasks-api/app/domain/entity"
	"github.com/JoyceEllenNeryTeles/test/tasks-api/app/domain/repository"
	"github.com/gin-gonic/gin"
)

type TaskUsecase struct {
	repository repository.TaskRepository
}

func NewTaskUsecase(repository repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		repository: repository,
	}
}

func (tuc *TaskUsecase) GetTasks() ([]entity.Task, error) {
	return tuc.repository.GetTasks()
}

func (tuc *TaskUsecase) AddTask(ctx *gin.Context, task entity.Task) (int64, error) {
	id, err := tuc.repository.AddTask(task)
	return id, err
}

func (tuc *TaskUsecase) FindTaskById(id int64) (*entity.Task, error) {
	task, err := tuc.repository.FindTaskById(id)
	return task, err
}

func (tuc *TaskUsecase) FindTaskByOwner(owner string) ([]entity.Task, error) {
	tasks, err := tuc.repository.FindTasksByOwner(owner)
	return tasks, err
}

func (tuc *TaskUsecase) UpdateTask(ctx *gin.Context, task entity.Task) error {
	err := tuc.repository.UpdateTask(task)
	return err
}

func (tuc *TaskUsecase) DeleteTask(id int64) error {
	err := tuc.repository.DeleteTask(id)
	return err
}

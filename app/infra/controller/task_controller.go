package controller

import (
	"net/http"
	"strconv"

	"github.com/JoyceEllenNeryTeles/tasks-api/app/domain/entity"
	"github.com/JoyceEllenNeryTeles/tasks-api/app/interfaces"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskInterface interfaces.TaskUsecaseInterface
}

func NewTaskController(taskInterface interfaces.TaskUsecaseInterface) *TaskController {
	return &TaskController{
		taskInterface: taskInterface,
	}
}

func (c *TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := c.taskInterface.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) AddTask(ctx *gin.Context) {

	var task entity.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
		return
	}

	id, err := c.taskInterface.AddTask(ctx, task)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add task"})
		return
	}
	ctx.JSON(http.StatusCreated, id)
}

func (c *TaskController) FindTaskById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task ID"})
		return
	}

	task, err := c.taskInterface.FindTaskById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find task"})
		return
	}

	if task == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	var task entity.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
		return
	}

	if task.Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required for update"})
		return
	}

	err := c.taskInterface.UpdateTask(ctx, task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task ID"})
		return
	}

	err = c.taskInterface.DeleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

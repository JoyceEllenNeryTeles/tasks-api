package main

import (
	"github.com/JoyceEllenNeryTeles/tasks-api/app/domain/repository"
	"github.com/JoyceEllenNeryTeles/tasks-api/app/infra/controller"
	"github.com/JoyceEllenNeryTeles/tasks-api/app/infra/db"
	"github.com/JoyceEllenNeryTeles/tasks-api/app/interfaces"
	"github.com/JoyceEllenNeryTeles/tasks-api/app/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}

	taskRepository := repository.NewTaskRepository(dbConnection)
	taskUsecase := usecase.NewTaskUsecase(*taskRepository)
	taskInterface := interfaces.TaskUsecaseInterface(taskUsecase)
	taskController := controller.NewTaskController(taskInterface)

	server.POST("/tasks", taskController.AddTask)
	server.GET("/tasks", taskController.GetTasks)
	server.GET("/tasks/:id", taskController.FindTaskById)
	server.PUT("/tasks/:id", taskController.UpdateTask)
	server.DELETE("/tasks/:id", taskController.DeleteTask)

	server.Run(":8080")
}

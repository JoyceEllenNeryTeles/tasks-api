package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JoyceEllenNeryTeles/tasks-api/app/domain/entity"
	"github.com/JoyceEllenNeryTeles/tasks-api/app/infra/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockTaskUsecase struct {
	GetTasksFn     func() ([]entity.Task, error)
	AddTaskFn      func(ctx *gin.Context, task entity.Task) (int64, error)
	FindTaskByIdFn func(id int64) (*entity.Task, error)
	UpdateTaskFn   func(ctx *gin.Context, task entity.Task) error
	DeleteTaskFn   func(id int64) error
}

func (m *MockTaskUsecase) GetTasks() ([]entity.Task, error) {
	return m.GetTasksFn()
}
func (m *MockTaskUsecase) AddTask(ctx *gin.Context, task entity.Task) (int64, error) {
	return m.AddTaskFn(ctx, task)
}
func (m *MockTaskUsecase) FindTaskById(id int64) (*entity.Task, error) {
	return m.FindTaskByIdFn(id)
}
func (m *MockTaskUsecase) UpdateTask(ctx *gin.Context, task entity.Task) error {
	return m.UpdateTaskFn(ctx, task)
}
func (m *MockTaskUsecase) DeleteTask(id int64) error {
	return m.DeleteTaskFn(id)
}

func setupRouter(ctrl *controller.TaskController) *gin.Engine {
	r := gin.Default()
	r.GET("/tasks", ctrl.GetTasks)
	r.POST("/tasks", ctrl.AddTask)
	r.GET("/tasks/:id", ctrl.FindTaskById)
	r.PUT("/tasks", ctrl.UpdateTask)
	r.DELETE("/tasks/:id", ctrl.DeleteTask)
	return r
}

func TestGetTasksControllerSuccess(t *testing.T) {
	mock := &MockTaskUsecase{
		GetTasksFn: func() ([]entity.Task, error) {
			return []entity.Task{{Id: 1, Description: "Test"}}, nil
		},
	}
	ctrl := controller.NewTaskController(mock)
	router := setupRouter(ctrl)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var tasks []entity.Task
	err := json.Unmarshal(w.Body.Bytes(), &tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.Equal(t, int64(1), tasks[0].Id)
}

func TestGetTasksError(t *testing.T) {
	mock := &MockTaskUsecase{
		GetTasksFn: func() ([]entity.Task, error) {
			return nil, errors.New("fail")
		},
	}
	ctrl := controller.NewTaskController(mock)
	router := setupRouter(ctrl)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestAddTaskSuccess(t *testing.T) {
	mock := &MockTaskUsecase{
		AddTaskFn: func(ctx *gin.Context, task entity.Task) (int64, error) {
			return 42, nil
		},
	}
	ctrl := controller.NewTaskController(mock)
	router := setupRouter(ctrl)

	task := entity.Task{Description: "New Task"}
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var id int64
	json.Unmarshal(w.Body.Bytes(), &id)
	assert.Equal(t, int64(42), id)
}

func TestAddTaskBadRequest(t *testing.T) {
	mock := &MockTaskUsecase{}
	ctrl := controller.NewTaskController(mock)
	router := setupRouter(ctrl)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddTaskError(t *testing.T) {
	mock := &MockTaskUsecase{
		AddTaskFn: func(ctx *gin.Context, task entity.Task) (int64, error) {
			return 0, errors.New("fail")
		},
	}
	ctrl := controller.NewTaskController(mock)
	router := setupRouter(ctrl)

	task := entity.Task{Description: "New Task"}
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFindTaskByIdSuccess(t *testing.T) {
	mock := &MockTaskUsecase{
		FindTaskByIdFn: func(id int64) (*entity.Task, error) {
			return &entity.Task{Id: id, Description: "Found"}, nil
		},
	}
	ctrl := controller.NewTaskController(mock)
	router := setupRouter(ctrl)

	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var task entity.Task
	json.Unmarshal(w.Body.Bytes(), &task)
	assert.Equal(t, int64(1), task.Id)
}

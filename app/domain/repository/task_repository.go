package repository

import (
	"database/sql"
	"fmt"

	"github.com/JoyceEllenNeryTeles/tasks-api/app/domain/entity"
)

type TaskRepository struct {
	coonnection *sql.DB
}

func NewTaskRepository(connection *sql.DB) *TaskRepository {
	return &TaskRepository{
		coonnection: connection,
	}
}

func (tr *TaskRepository) GetTasks() ([]entity.Task, error) {
	var tasks []entity.Task
	rows, err := tr.coonnection.Query("SELECT id, description FROM tasks_table order by id")
	if err != nil {
		fmt.Errorf("Error querying tasks: %v", err)
		return []entity.Task{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.Id, &task.Description); err != nil {
			fmt.Errorf("Error scanning tasks: %v", err)
			return []entity.Task{}, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("Error whiler retrieving tasks: %v", err)
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (tr *TaskRepository) AddTask(task entity.Task) (int64, error) {
	query, err := tr.coonnection.Prepare("INSERT INTO tasks_table (description) VALUES ($1) returning id")
	if err != nil {
		fmt.Errorf("Error preparinf insert task: %v", err)
		return 0, err
	}

	err = query.QueryRow(task.Description).Scan(&task.Id)
	if err != nil {
		fmt.Errorf("Error scanning inserted task ID: %v", err)
		return 0, err
	}
	query.Close()
	return int64(task.Id), nil
}

func (tr *TaskRepository) FindTaskById(id int64) (*entity.Task, error) {
	query := "SELECT id, description FROM tasks_table WHERE id = $1"
	row := tr.coonnection.QueryRow(query, id)
	var task entity.Task
	err := row.Scan(&task.Id, &task.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Task not found
		}
		fmt.Errorf("Error scanning task by ID: %v", err)
		return nil, err
	}
	return &task, nil
}

func (tr *TaskRepository) UpdateTask(task entity.Task) error {
	query := "UPDATE tasks_table SET description = $1 WHERE id = $2"
	_, err := tr.coonnection.Exec(query, task.Description, task.Id)
	if err != nil {
		fmt.Errorf("Error updating task: %v", err)
		return err
	}
	return nil
}

func (tr *TaskRepository) DeleteTask(id int64) error {
	query := "DELETE FROM tasks_table WHERE id = $1"
	_, err := tr.coonnection.Exec(query, id)
	if err != nil {
		fmt.Errorf("Error deleting task: %v", err)
		return err
	}
	return nil
}

package repository

import (
	"database/sql"
	"fmt"

	"github.com/JoyceEllenNeryTeles/test/tasks-api/app/domain/entity"
)

type TaskRepository struct {
	connection *sql.DB
}

func NewTaskRepository(connection *sql.DB) *TaskRepository {
	return &TaskRepository{
		connection: connection,
	}
}

func (tr *TaskRepository) GetTasks() ([]entity.Task, error) {
	var tasks []entity.Task
	rows, err := tr.connection.Query("SELECT id, description, owner, status  FROM tasks_table order by id")
	if err != nil {
		fmt.Errorf("Error querying tasks: %v", err)
		return []entity.Task{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.Id, &task.Description, &task.Owner, &task.Status); err != nil {
			fmt.Errorf("Error scanning tasks: %v", err)
			return []entity.Task{}, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("Error while retrieving tasks: %v", err)
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (tr *TaskRepository) AddTask(task entity.Task) (int64, error) {
	query, err := tr.connection.Prepare("INSERT INTO tasks_table (description, owner, status) VALUES ($1, $2, $3) returning id")
	if err != nil {
		fmt.Errorf("Error preparing insert task: %v", err)
		return 0, err
	}

	err = query.QueryRow(task.Description, task.Owner, task.Status).Scan(&task.Id)
	if err != nil {
		fmt.Errorf("Error scanning inserted task ID: %v", err)
		return 0, err
	}
	query.Close()
	return int64(task.Id), nil
}

func (tr *TaskRepository) FindTaskById(id int64) (*entity.Task, error) {
	query := "SELECT id, description, owner, status FROM tasks_table WHERE id = $1"
	row := tr.connection.QueryRow(query, id)
	var task entity.Task
	err := row.Scan(&task.Id, &task.Description, &task.Owner, &task.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Errorf("Error scanning task by ID: %v", err)
		return nil, err
	}
	return &task, nil
}

func (tr *TaskRepository) FindTasksByOwner(owner string) ([]entity.Task, error) {
	query := "SELECT id, description, owner, status FROM tasks_table WHERE owner = $1"
	rows, err := tr.connection.Query(query, owner)
	if err != nil {
		fmt.Errorf("Error querying tasks by owner: %v", err)
		return []entity.Task{}, err
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.Id, &task.Description, &task.Owner, &task.Status); err != nil {
			fmt.Errorf("Error scanning tasks by owner: %v", err)
			return []entity.Task{}, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("Error while retrieving tasks by owner: %v", err)
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (tr *TaskRepository) UpdateTask(task entity.Task) error {
	query := "UPDATE tasks_table SET description = $1, owner = $2, status = $3 WHERE id = $4"
	_, err := tr.connection.Exec(query, task.Description, task.Owner, task.Status, task.Id)
	if err != nil {
		fmt.Errorf("Error updating task: %v", err)
		return err
	}
	return nil
}

func (tr *TaskRepository) DeleteTask(id int64) error {
	query := "DELETE FROM tasks_table WHERE id = $1"
	_, err := tr.connection.Exec(query, id)
	if err != nil {
		fmt.Errorf("Error deleting task: %v", err)
		return err
	}
	return nil
}

package db

import (
	"context"
	"database/sql"
	"errors"
	"golangTestCase/internal/app/models"
)

func CreateTask(ctx context.Context, task *models.Task) error {
	query := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	err := db.QueryRowContext(ctx, query, task.Title, task.Description, task.Status).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func GetTask(ctx context.Context, id int) (*models.Task, error) {
	task := &models.Task{}
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id=$1`
	err := db.QueryRowContext(ctx, query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return task, nil
}

func UpdateTask(ctx context.Context, task *models.Task) error {
	query := `UPDATE tasks SET title=$2, description=$3, status=$4, updated_at=NOW() WHERE id=$1 RETURNING updated_at`
	err := db.QueryRowContext(ctx, query, task.ID, task.Title, task.Description, task.Status).Scan(&task.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(ctx context.Context, id int) error {
	query := `DELETE FROM tasks WHERE id=$1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

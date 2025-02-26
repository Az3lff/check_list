package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/Az3lff/check_list/internal/config"
	"github.com/Az3lff/check_list/internal/models"
)

type TaskRepository struct {
	db *pgx.Conn
}

func NewTaskRepository(ctx context.Context, cfg config.Postgres) (*TaskRepository, error) {
	conn, err := pgx.Connect(ctx, cfg.Connection)
	if err != nil {
		return nil, err
	}

	return &TaskRepository{
		db: conn,
	}, nil
}

func (r *TaskRepository) CreateTask(ctx context.Context, req models.CreateTaskRequest) (*models.TaskIDResponse, error) {
	var resp models.TaskIDResponse
	if err := r.db.QueryRow(ctx, queryCreateTask,
		req.UserID,
		req.Title,
		req.Description,
		req.Done,
	).Scan(&resp.TaskID); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *TaskRepository) GetAllTasksByUserID(ctx context.Context, userID int64) (*models.GetListResponse, error) {
	rows, err := r.db.Query(ctx, queryGetAllTasksByUserID, userID)
	if err != nil {
		return nil, err
	}

	tasks, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Task])
	if err != nil {
		return nil, err
	}

	return &models.GetListResponse{
		List: tasks,
	}, nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, req models.DeleteTaskRequest) (*models.TaskIDResponse, error) {
	var resp models.TaskIDResponse
	if err := r.db.QueryRow(ctx, queryDeleteTask,
		req.TaskID,
		req.UserID,
	).Scan(&resp.TaskID); err != nil {
		return nil, fmt.Errorf("db error: %w", err)
	}

	return &resp, nil
}

func (r *TaskRepository) DoneTask(ctx context.Context, req models.DoneTaskRequest) (*models.TaskIDResponse, error) {
	var resp models.TaskIDResponse
	if err := r.db.QueryRow(ctx, queryDoneTask,
		req.Done,
		req.TaskID,
		req.UserID,
	).Scan(&resp.TaskID); err != nil {
		return nil, fmt.Errorf("db error: %w", err)
	}

	return &resp, nil
}

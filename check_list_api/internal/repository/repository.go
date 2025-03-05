package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Az3lff/check_list/internal/models"
	"github.com/Az3lff/check_list/internal/repository/check_list"
)

type Repositories struct {
	CheckList
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		check_list.NewCheckListRepo(db),
	}
}

type CheckList interface {
	CreateTask(ctx context.Context, req models.CreateTaskRequest) (*models.TaskIDResponse, error)
	GetTaskByID(ctx context.Context, req models.GetTaskRequest) (*models.Task, error)
	GetAllTasksByUserID(ctx context.Context, userID int64) (*models.GetListResponse, error)
	DeleteTask(ctx context.Context, req models.DeleteTaskRequest) (*models.TaskIDResponse, error)
	DoneTask(ctx context.Context, req models.DoneTaskRequest) (*models.TaskIDResponse, error)
}

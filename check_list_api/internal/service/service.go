package service

import (
	"context"

	"github.com/Az3lff/check_list/internal/models"
	"github.com/Az3lff/check_list/internal/repository"
	"github.com/Az3lff/check_list/internal/service/check_list"
)

type Services struct {
	CheckList
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		check_list.NewCheckListService(repos.CheckList),
	}
}

type CheckList interface {
	CreateTask(ctx context.Context, req models.CreateTaskRequest) (*models.TaskIDResponse, error)
	GetTask(ctx context.Context, req models.GetTaskRequest) (*models.Task, error)
	GetList(ctx context.Context, userID int64) (*models.GetListResponse, error)
	DeleteTask(ctx context.Context, req models.DeleteTaskRequest) (*models.TaskIDResponse, error)
	DoneTask(ctx context.Context, req models.DoneTaskRequest) (*models.TaskIDResponse, error)
}

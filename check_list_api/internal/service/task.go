package service

import (
	"context"

	"github.com/Az3lff/check_list/internal/models"
	"github.com/Az3lff/check_list/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repository *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repository,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, req models.CreateTaskRequest) (*models.TaskIDResponse, error) {
	resp, err := s.repo.CreateTask(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *TaskService) GetList(ctx context.Context, userID int64) (*models.GetListResponse, error) {
	resp, err := s.repo.GetAllTasksByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, req models.DeleteTaskRequest) (*models.TaskIDResponse, error) {
	resp, err := s.repo.DeleteTask(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *TaskService) DoneTask(ctx context.Context, req models.DoneTaskRequest) (*models.TaskIDResponse, error) {
	resp, err := s.repo.DoneTask(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

package check_list

import (
	"context"

	"github.com/Az3lff/check_list/internal/models"
	"github.com/Az3lff/check_list/internal/repository"
)

type CheckListService struct {
	repo repository.CheckList
}

func NewCheckListService(repo repository.CheckList) *CheckListService {
	return &CheckListService{
		repo: repo,
	}
}

func (s *CheckListService) CreateTask(ctx context.Context, req models.CreateTaskRequest) (*models.TaskIDResponse, error) {
	resp, err := s.repo.CreateTask(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CheckListService) GetTask(ctx context.Context, req models.GetTaskRequest) (*models.Task, error) {
	resp, err := s.repo.GetTask(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CheckListService) GetList(ctx context.Context, userID int64) (*models.GetListResponse, error) {
	resp, err := s.repo.GetList(ctx, userID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CheckListService) DeleteTask(ctx context.Context, req models.DeleteTaskRequest) (*models.TaskIDResponse, error) {
	resp, err := s.repo.DeleteTask(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CheckListService) DoneTask(ctx context.Context, req models.DoneTaskRequest) (*models.TaskIDResponse, error) {
	resp, err := s.repo.DoneTask(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

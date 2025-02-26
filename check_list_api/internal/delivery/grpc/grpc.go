package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Az3lff/check_list/internal/delivery/grpc/task"
	"github.com/Az3lff/check_list/internal/models"
	"github.com/Az3lff/check_list/internal/service"
)

type server struct {
	task.UnimplementedTaskServer

	svc *service.TaskService
}

func Register(gRPC *grpc.Server, service *service.TaskService) {
	task.RegisterTaskServer(gRPC, &server{
		svc: service,
	})
}

func (s *server) CreateTask(ctx context.Context, req *task.CreateTaskRequest) (*task.TaskIDResponse, error) {
	resp, err := s.svc.CreateTask(ctx, models.CreateTaskRequest{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Done:        req.Done,
	})
	if err != nil {
		return nil, err
	}

	return &task.TaskIDResponse{
		TaskID: resp.TaskID,
	}, nil
}

func (s *server) GetList(ctx context.Context, req *task.GetListRequest) (*task.GetListResponse, error) {
	resp, err := s.svc.GetList(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	var list task.GetListResponse
	if resp.List != nil {
		list.Tasks = make([]*task.GetListResponse_Task, 0, len(resp.List))

		for i := range resp.List {
			list.Tasks = append(list.Tasks, &task.GetListResponse_Task{
				TaskID:      resp.List[i].TaskID,
				Title:       resp.List[i].Title,
				Description: resp.List[i].Description,
				Done:        resp.List[i].Done,
			})
		}
	}

	return &list, nil
}

func (s *server) DeleteTask(ctx context.Context, req *task.DeleteTaskRequest) (*task.TaskIDResponse, error) {
	resp, err := s.svc.DeleteTask(ctx, models.DeleteTaskRequest{
		UserID: req.UserID,
		TaskID: req.TaskID,
	})
	if err != nil {
		return nil, err
	}

	return &task.TaskIDResponse{
		TaskID: resp.TaskID,
	}, nil
}

func (s *server) DoneTask(ctx context.Context, req *task.DoneTaskRequest) (*task.TaskIDResponse, error) {
	resp, err := s.svc.DoneTask(ctx, models.DoneTaskRequest{
		UserID: req.UserID,
		TaskID: req.TaskID,
	})
	if err != nil {
		return nil, err
	}

	return &task.TaskIDResponse{
		TaskID: resp.TaskID,
	}, nil
}

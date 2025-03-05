package check_list

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Az3lff/check_list/internal/models"
)

type CheckListRepo struct {
	db *pgxpool.Pool
}

func NewCheckListRepo(db *pgxpool.Pool) *CheckListRepo {
	return &CheckListRepo{
		db: db,
	}
}

func (r *CheckListRepo) CreateTask(ctx context.Context, req models.CreateTaskRequest) (*models.TaskIDResponse, error) {
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

func (r *CheckListRepo) GetTaskByID(ctx context.Context, req models.GetTaskRequest) (*models.Task, error) {
	rows, err := r.db.Query(ctx, queryGetTaskByUserIDAndTaskID, req.UserID, req.TaskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resp models.Task
	if err := rows.Scan(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *CheckListRepo) GetAllTasksByUserID(ctx context.Context, userID int64) (*models.GetListResponse, error) {
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

func (r *CheckListRepo) DeleteTask(ctx context.Context, req models.DeleteTaskRequest) (*models.TaskIDResponse, error) {
	var resp models.TaskIDResponse
	if err := r.db.QueryRow(ctx, queryDeleteTask,
		req.TaskID,
		req.UserID,
	).Scan(&resp.TaskID); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *CheckListRepo) DoneTask(ctx context.Context, req models.DoneTaskRequest) (*models.TaskIDResponse, error) {
	var resp models.TaskIDResponse
	if err := r.db.QueryRow(ctx, queryDoneTask,
		req.Done,
		req.TaskID,
		req.UserID,
	).Scan(&resp.TaskID); err != nil {
		return nil, err
	}

	return &resp, nil
}

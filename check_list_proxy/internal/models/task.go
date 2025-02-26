package models

type Task struct {
	TaskID      int64  `json:"task_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description" validate:"required"`
	Done        *bool  `json:"done" validate:"required"`
}

type DeleteTaskRequest struct {
	TaskID int64 `json:"task_id" validate:"required,gt=0"`
}

type DoneTaskRequest struct {
	TaskID int64 `json:"task_id" validate:"required,gt=0"`
	Done   *bool `json:"done" validate:"required"`
}

type TaskIDResponse struct {
	TaskID int64 `json:"task_id"`
}

type GetListResponse struct {
	List []Task `json:"list"`
}

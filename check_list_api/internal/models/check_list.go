package models

type Task struct {
	TaskID      int64
	Title       string
	Description string
	Done        bool
}

type CreateTaskRequest struct {
	UserID      int64
	Title       string
	Description string
	Done        bool
}

type GetTaskRequest struct {
	UserID int64
	TaskID int64
}

type DeleteTaskRequest struct {
	UserID int64
	TaskID int64
}

type DoneTaskRequest struct {
	UserID int64
	TaskID int64
	Done   bool
}

type TaskIDResponse struct {
	TaskID int64
}

type GetListResponse struct {
	List []Task
}

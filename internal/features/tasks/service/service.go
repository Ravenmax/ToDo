package tasks_service

import (
	"context"

	"github.com/Ravenmax/ToDo/internal/core/domain"
)

type TasksService struct {
	tasksRepository TasksRepository
}
type TasksRepository interface {
	Createtask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)
	GetTasks(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]domain.Task, error)
	GetTask(
		ctx context.Context,
		taskid int,
	) (domain.Task, error)
	DeleteTask(
		ctx context.Context,
		taskid int,
	) error
	PatchTask(
		ctx context.Context,
		taskid int,
		taskPatched domain.Task,
	) (domain.Task, error)
}

func NewTasksService(
	tasksRepository TasksRepository,
) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}

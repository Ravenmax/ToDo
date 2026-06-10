package tasks_service

import (
	"context"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	"github.com/google/uuid"
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
		userID *uuid.UUID,
		limit *int,
		offset *int,
	) ([]domain.Task, error)
	GetTask(
		ctx context.Context,
		taskid uuid.UUID,
	) (domain.Task, error)
	DeleteTask(
		ctx context.Context,
		taskid uuid.UUID,
	) error
	PatchTask(
		ctx context.Context,
		taskid uuid.UUID,
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

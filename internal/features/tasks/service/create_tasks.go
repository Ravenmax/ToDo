package tasks_service

import (
	"context"
	"fmt"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	"github.com/google/uuid"
)

func (s *TasksService) CreateTask(
	ctx context.Context,
	title string,
	description *string,
	authorUserID uuid.UUID,
) (domain.Task, error) {
	task := domain.CreateTask(
		title,
		description,
		authorUserID,
	)
	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate task domain: %w", err)
	}
	task, err := s.tasksRepository.Createtask(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("create task: %w", err)
	}
	return task, nil
}

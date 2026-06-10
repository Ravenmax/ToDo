package tasks_service

import (
	"context"
	"fmt"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	"github.com/google/uuid"
)

func (s *TasksService) PatchTask(
	ctx context.Context,
	taskid uuid.UUID,
	patch domain.TaskPatch,
) (domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, taskid)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}
	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("patch task: %w", err)
	}
	taskPatched, err := s.tasksRepository.PatchTask(ctx, taskid, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patch user: %w", err)
	}
	return taskPatched, nil

}

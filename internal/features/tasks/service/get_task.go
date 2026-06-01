package tasks_service

import (
	"context"
	"fmt"

	"github.com/Ravenmax/ToDo/internal/core/domain"
)

func (s *TasksService) GetTask(
	ctx context.Context,
	taskid int,
) (domain.Task, error) {

	task, err := s.tasksRepository.GetTask(ctx, taskid)
	if err != nil {
		return domain.Task{}, fmt.Errorf("cant get task from repository: %w", err)
	}
	return task, nil
}

package tasks_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *TasksService) DeleteTask(
	ctx context.Context,
	taskid uuid.UUID,
) error {
	err := s.tasksRepository.DeleteTask(ctx, taskid)
	if err != nil {
		return fmt.Errorf("failed to delete from repository: %w", err)
	}
	return nil
}

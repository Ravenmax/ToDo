package tasks_service

import (
	"context"
	"fmt"
)

func (s *TasksService) DeleteTask(
	ctx context.Context,
	taskid int,
) error {
	err := s.tasksRepository.DeleteTask(ctx, taskid)
	if err != nil {
		return fmt.Errorf("failed to delete from repository: %w", err)
	}
	return nil
}

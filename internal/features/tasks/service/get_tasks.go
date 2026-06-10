package tasks_service

import (
	"context"
	"fmt"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_errors "github.com/Ravenmax/ToDo/internal/core/errors"
	"github.com/google/uuid"
)

func (s *TasksService) GetTasks(
	ctx context.Context,
	userID *uuid.UUID,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	taskDomains, err := s.tasksRepository.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get task from repository: %w", err)
	}
	return taskDomains, nil
}

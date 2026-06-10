package tasks_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/Ravenmax/ToDo/internal/core/errors"
	"github.com/google/uuid"
)

func (r *TasksRepository) DeleteTask(
	ctx context.Context,
	taskid uuid.UUID,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM todoapp.tasks
	WHERE ID=$1
	`
	cmdTag, err := r.pool.Exec(ctx, query, taskid)

	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("task with id=%d : %w", taskid, core_errors.ErrNotFound)
	}

	return nil
}

package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_errors "github.com/Ravenmax/ToDo/internal/core/errors"
	core_postgres_pool "github.com/Ravenmax/ToDo/internal/core/repository/postgres/pull"
)

func (r *TasksRepository) PatchTask(
	ctx context.Context,
	taskid int,
	taskPatched domain.Task,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE todoapp.tasks
	SET
		title = $1,
		description = $2,
		completed = $3,
		completed_at = $4,
		version = version+1
	WHERE id=$5 AND version=$6	
	RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id	
	`
	row := r.pool.QueryRow(
		ctx,
		query,
		taskPatched.Title,
		taskPatched.Description,
		taskPatched.Completed,
		taskPatched.CompletedAt,
		taskid,
		taskPatched.Version,
	)
	var taskModel TaskModel
	err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorUserId,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("user with ID: %d concurency contest: %w", taskid, core_errors.ErrConflict)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	taskDomain := taskDomainFromModel(taskModel)

	return taskDomain, nil
}

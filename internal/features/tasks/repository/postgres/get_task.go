package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_errors "github.com/Ravenmax/ToDo/internal/core/errors"
	core_postgres_pool "github.com/Ravenmax/ToDo/internal/core/repository/postgres/pull"
)

func (r *TasksRepository) GetTask(
	ctx context.Context,
	taskid int,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks
	WHERE id=$1
	ORDER BY id ASC
	`
	row := r.pool.QueryRow(ctx, query, taskid)
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
			return domain.Task{}, fmt.Errorf(
				"task with id=%d: %w",
				taskid,
				core_errors.ErrNotFound,
			)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	taskDomain := domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserId,
	)

	return taskDomain, nil
}

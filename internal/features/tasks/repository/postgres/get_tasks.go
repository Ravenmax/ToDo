package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_errors "github.com/Ravenmax/ToDo/internal/core/errors"
	core_postgres_pool "github.com/Ravenmax/ToDo/internal/core/repository/postgres/pull"
)

func (r *TasksRepository) GetTasks(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	var query string

	query = `
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks
	
	%s
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;`

	args := []any{limit, offset}

	if userID != nil {
		query = fmt.Sprintf(query, "WHERE author_user_id = $3")
		args = append(args, userID)
	} else {
		query = fmt.Sprintf(query, "")
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("select tasks: %w", err)
	}
	//обязательно закрываем rows.Close
	defer rows.Close()

	var taskModels []TaskModel
	for rows.Next() {
		var taskModel TaskModel
		err := rows.Scan(
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
			if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
				return nil,
					fmt.Errorf(
						"%v, user withID=%d not found: %w",
						err,
						userID,
						core_errors.ErrNotFound,
					)
			}
			return nil, fmt.Errorf("scan error: %w", err)
		}
		taskModels = append(taskModels, taskModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	taskDomains := taskDomainsFromModels(taskModels)

	return taskDomains, nil
}

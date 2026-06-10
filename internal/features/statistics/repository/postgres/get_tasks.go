package statistics_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_errors "github.com/Ravenmax/ToDo/internal/core/errors"
	core_postgres_pool "github.com/Ravenmax/ToDo/internal/core/repository/postgres/pull"
	"github.com/google/uuid"
)

func (r *StatisticsRepository) GetTasks(
	ctx context.Context,
	userID *uuid.UUID,
	from *time.Time,
	to *time.Time,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks
	`)
	args := []any{}
	conditions := []string{}
	if userID != nil {
		conditions = append(conditions, fmt.Sprintf("author_user_id=$%d", len(args)+1))
		args = append(args, userID)
	}
	if from != nil {
		conditions = append(conditions, fmt.Sprintf("created_at>=$%d", len(args)+1))
		args = append(args, from)
	}
	if to != nil {
		conditions = append(conditions, fmt.Sprintf("created_at<$%d", len(args)+1))
		args = append(args, to)
	}
	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE " + strings.Join(conditions, " AND "))
	}

	queryBuilder.WriteString(" ORDER BY id ASC")
	rows, err := r.pool.Query(ctx, queryBuilder.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("select tasks: %w", err)
	}
	defer rows.Close()

	var taskModels []TaskModel
	for rows.Next() {
		var taskModel TaskModel
		if err := taskModel.Scan(rows); err != nil {
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

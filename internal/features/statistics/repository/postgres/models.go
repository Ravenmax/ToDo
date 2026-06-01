package statistics_postgres_repository

import (
	"time"

	"github.com/Ravenmax/ToDo/internal/core/domain"
)

type TaskModel struct {
	ID           int        `db:"id"`
	Version      int64      `db:"version"`
	Title        string     `db:"title"`
	Description  *string    `db:"decription"`
	Completed    bool       `db:"completed"`
	CreatedAt    time.Time  `db:"created_at"`
	CompletedAt  *time.Time `db:"completed_at"`
	AuthorUserId int        `db:"author_user_id"`
}

func taskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserId,
	)
}

func taskDomainsFromModels(tasks []TaskModel) []domain.Task {
	taskDomains := make([]domain.Task, len(tasks))
	for i, task := range tasks {
		taskDomains[i] = taskDomainFromModel(task)
	}
	return taskDomains
}

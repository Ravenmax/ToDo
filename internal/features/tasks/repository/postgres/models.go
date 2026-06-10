package tasks_postgres_repository

import (
	"time"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_postgres_pool "github.com/Ravenmax/ToDo/internal/core/repository/postgres/pull"
	"github.com/google/uuid"
)

type TaskModel struct {
	ID           uuid.UUID  `db:"id"`
	Version      int        `db:"version"`
	Title        string     `db:"title"`
	Description  *string    `db:"decription"`
	Completed    bool       `db:"completed"`
	CreatedAt    time.Time  `db:"created_at"`
	CompletedAt  *time.Time `db:"completed_at"`
	AuthorUserID uuid.UUID  `db:"author_user_id"`
}

func (m *TaskModel) Scan(row core_postgres_pool.Row) error {
	return row.Scan(
		&m.ID,
		&m.Version,
		&m.Title,
		&m.Description,
		&m.Completed,
		&m.CreatedAt,
		&m.CompletedAt,
		&m.AuthorUserID,
	)
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
		taskModel.AuthorUserID,
	)
}

func taskDomainsFromModels(tasks []TaskModel) []domain.Task {
	taskDomains := make([]domain.Task, len(tasks))
	for i, task := range tasks {
		taskDomains[i] = taskDomainFromModel(task)
	}
	return taskDomains
}

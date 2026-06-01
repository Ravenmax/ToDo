package tasks_transport_http

import (
	"time"

	"github.com/Ravenmax/ToDo/internal/core/domain"
)

type TaskDTOResponce struct {
	ID           int        `json:"id"`
	Version      int64      `json:"version"`
	Title        string     `json:"title"`
	Desctiption  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"createdAt"`
	CompletedAt  *time.Time `json:"completedAt"`
	AuthorUserID int        `json:"authorUserID"`
}

func TaskDTOFromDomain(task domain.Task) TaskDTOResponce {
	return TaskDTOResponce{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Desctiption:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}
func TasksDTOFromDomains(tasks []domain.Task) []TaskDTOResponce {
	taskDTO := make([]TaskDTOResponce, len(tasks))
	for i, task := range tasks {
		taskDTO[i] = TaskDTOFromDomain(task)
	}
	return taskDTO
}

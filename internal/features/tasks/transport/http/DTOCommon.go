package tasks_transport_http

import (
	"time"

	"github.com/Ravenmax/ToDo/internal/core/domain"
)

type TaskDTOResponce struct {
	ID           int        `json:"id"                  example:"10"`
	Version      int64      `json:"version"             example:"3"`
	Title        string     `json:"title"               example:"Помыть посуду"`
	Desctiption  *string    `json:"description"         example:"срочно"`
	Completed    bool       `json:"completed"           example:"true"`
	CreatedAt    time.Time  `json:"createdAt"           example:"10.02.2026"`
	CompletedAt  *time.Time `json:"completedAt"         example:"11.02.2026"`
	AuthorUserID int        `json:"authorUserID"        example:"2"`
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

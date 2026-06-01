package domain

import "time"

type Statistics struct {
	TasksCreated               int
	TasksCompleted             int
	TasksCompletedRate         *float64
	TasksAverageCompletionTIme *time.Duration
}

func NewStatistics(
	taskCreated int,
	taskComleted int,
	taskCompletedRate *float64,
	taskAverageCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		TasksCreated:               taskComleted,
		TasksCompleted:             taskComleted,
		TasksCompletedRate:         taskCompletedRate,
		TasksAverageCompletionTIme: taskAverageCompletionTime,
	}
}

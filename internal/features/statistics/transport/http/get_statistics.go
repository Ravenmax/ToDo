package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_request "github.com/Ravenmax/ToDo/internal/core/transport/http/request"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
)

type QueryParams struct {
	userId *int
	from   *time.Time
	to     *time.Time
}
type GetStatisticsResponse struct {
	TaskCreated               int      `json:"task_created"`
	TaskCompleted             int      `json:"task_completed"`
	TaskCompletedRate         *float64 `json:"task_completed_rate"`
	TaskAverageCompletionTIme *string  `json:"task_average_completed_time"`
}

func toDTOFromDomain(statistics domain.Statistics) GetStatisticsResponse {
	var avgTime *string
	if statistics.TasksAverageCompletionTIme != nil {
		duration := statistics.TasksAverageCompletionTIme.String()
		avgTime = &duration
	}

	return GetStatisticsResponse{
		TaskCreated:               statistics.TasksCreated,
		TaskCompleted:             statistics.TasksCompleted,
		TaskCompletedRate:         statistics.TasksCompletedRate,
		TaskAverageCompletionTIme: avgTime,
	}
}
func (h *StattisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	queryParams, err := GetUserIDFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID/from/to query params")
		return
	}
	statisticsDomain, err := h.statisticsService.GetStatistics(ctx, queryParams.userId, queryParams.from, queryParams.to)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed get statistics")
	}
	response := toDTOFromDomain(statisticsDomain)

	responseHandler.JSONResponse(response, http.StatusOK)
}

func GetUserIDFromToQueryParams(r *http.Request) (QueryParams, error) {
	const (
		userIDQueryParam = "user_id"
		fromQueryParam   = "from"
		toQueryParam     = "to"
	)
	var (
		resultParams QueryParams
		err          error
	)
	resultParams.userId, err = core_http_request.GetIntQueryParam(r, userIDQueryParam)
	if err != nil {
		return QueryParams{}, fmt.Errorf("get `userID` params: %w", err)
	}
	resultParams.from, err = core_http_request.GetDateQueryParam(r, fromQueryParam)
	if err != nil {
		return QueryParams{}, fmt.Errorf("get `from` params: %w", err)
	}
	resultParams.to, err = core_http_request.GetDateQueryParam(r, toQueryParam)
	if err != nil {
		return QueryParams{}, fmt.Errorf("get `to` params: %w", err)
	}

	return resultParams, nil
}

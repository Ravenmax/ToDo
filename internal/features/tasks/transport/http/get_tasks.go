package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_request "github.com/Ravenmax/ToDo/internal/core/transport/http/request"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDTOResponce

type TasksQueryParams struct {
	userid *int
	limit  *int
	offset *int
}

func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	queryParams, err := getUserIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userid/limit/offset query params")
		return
	}
	taskDomains, err := h.tasksService.GetTasks(ctx, queryParams.userid, queryParams.limit, queryParams.offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks")
	}
	response := GetTasksResponse(TasksDTOFromDomains(taskDomains))

	responseHandler.JSONResponse(response, http.StatusOK)

}

func getUserIDLimitOffsetQueryParams(r *http.Request) (TasksQueryParams, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
		userIDQueryParamKey = "user_id"
	)
	var (
		resultQueryParams TasksQueryParams
		err               error
	)
	resultQueryParams.limit, err = core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return TasksQueryParams{}, fmt.Errorf("get 'limit' query param: %w", err)
	}
	resultQueryParams.offset, err = core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return TasksQueryParams{}, fmt.Errorf("get 'offset' query param: %w", err)
	}
	resultQueryParams.userid, err = core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return TasksQueryParams{}, fmt.Errorf("get 'limit' query param: %w", err)
	}

	return resultQueryParams, nil
}

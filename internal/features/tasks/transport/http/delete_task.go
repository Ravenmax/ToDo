package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_request "github.com/Ravenmax/ToDo/internal/core/transport/http/request"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
)

func (h *TasksHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskid, err := core_http_request.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "cant get id path from request")
	}

	err = h.tasksService.DeleteTask(ctx, taskid)
	if err != nil {
		responseHandler.ErrorResponse(err, fmt.Sprintf("failde to delete task with id=%d", taskid))
	}
	responseHandler.NoContentResponse()
}

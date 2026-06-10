package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_request "github.com/Ravenmax/ToDo/internal/core/transport/http/request"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
)

// DeleteTask    godoc
// @Summary      Удаление задачи
// @Description  Удаление задачи по входящему ID
// @Tags         Tasks
// @Param        id path int true "ID удаляемой задачи"
// @Success      204  "Успешное удаление задачи"
// @Failure      400  {object}  core_http_response.ErrorResponse "Bad request"
// @Failure      404  {object}  core_http_response.ErrorResponse "Task not found"
// @Failure      500  {object}  core_http_response.ErrorResponse "Internal server error"
// @Router		 /tasks/{id} [delete]
func (h *TasksHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskid, err := core_http_request.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "cant get id path from request")
	}

	err = h.tasksService.DeleteTask(ctx, taskid)
	if err != nil {
		responseHandler.ErrorResponse(err, fmt.Sprintf("failde to delete task with id=%d", taskid))
	}
	responseHandler.NoContentResponse()
}

package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_request "github.com/Ravenmax/ToDo/internal/core/transport/http/request"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
)

// GetTask       godoc
// @Summary      Получение задачи
// @Description  Получения задачи по ID
// @Tags         Tasks
// @Produce      json
// @Param        id path int true "ID задачи"
// @Success      200  {object}  CreateTaskResponse "Найденный пользователь"
// @Failure      400  {object}  core_http_response.ErrorResponse "Bad request"
// @Failure      404  {object}  core_http_response.ErrorResponse "Task not found"
// @Failure      500  {object}  core_http_response.ErrorResponse "Internal server error"
// @Router		 /tasks/{id} [get]
func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get TaskID path Value")
		return
	}

	taskDomain, err := h.tasksService.GetTask(ctx, taskID)
	if err != nil {
		responseHandler.ErrorResponse(err, fmt.Sprintf("failed to get task with ID=%d", taskID))
		return
	}
	response := CreateTaskResponse(TaskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)

}

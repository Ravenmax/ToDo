package tasks_transport_http

import (
	"net/http"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_request "github.com/Ravenmax/ToDo/internal/core/transport/http/request"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Title        string    `json:"title" validate:"required,min=1,max=100"            example:"Сделать дз"`
	Description  *string   `json:"description" validate:"omitempty,min=1,max=1000"    example:"Сделать домашние задание по математике"`
	AuthorUserID uuid.UUID `json:"author_user_id" validate:"required"                 example:"2"`
}

type CreateTaskResponse TaskDTOResponce

// CreateTask  godoc
// @Summary      Создание задачи
// @Description  Создание задачи в системе из json body запроса
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Param        request body CreateTaskRequest true "CreateTask тело запроса"
// @Success      201  {object}  CreateTaskResponse "Успешно созданный пользователь"
// @Failure      400  {object}  core_http_response.ErrorResponse "Bad request"
// @Failure      500  {object}  core_http_response.ErrorResponse "Internal server error"
// @Router		 /tasks [post]
func (h *TasksHTTPHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	var request CreateTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate request",
		)
		return
	}
	taskDomain := domain.CreateTask(
		request.Title,
		request.Description,
		request.AuthorUserID,
	)
	taskDomain, err := h.tasksService.CreateTask(
		ctx,
		request.Title,
		request.Description,
		request.AuthorUserID,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create task",
		)
		return
	}
	response := TaskDTOFromDomain(taskDomain)

	responseHandler.JSONResponse(response, http.StatusOK)

}

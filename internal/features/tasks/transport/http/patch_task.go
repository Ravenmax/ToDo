package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_request "github.com/Ravenmax/ToDo/internal/core/transport/http/request"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
	core_http_types "github.com/Ravenmax/ToDo/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Comleted    core_http_types.Nullable[bool]   `json:"completed"`
}
type PatchTaskReponse TaskDTOResponce

func (r *PatchTaskRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("Title can't be null")
		}
		titleLen := len([]rune(*r.Title.Value))
		if titleLen < 1 || titleLen > 100 {
			return fmt.Errorf("Title must be beetwen 1 and 100 symbols")
		}
	}
	if r.Description.Set {
		descriptionLen := len([]rune(*r.Description.Value))
		if descriptionLen < 1 || descriptionLen > 1000 {
			return fmt.Errorf("description must be beetwen 1 and 100 symbols")
		}
	}
	if r.Comleted.Set {
		if r.Comleted.Value == nil {
			return fmt.Errorf("Completed can't be nil")
		}
	}
	return nil
}
func (h *TasksHTTPHandler) PatchTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"Failed to get taskID path value",
		)
		return
	}
	var request PatchTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate http request",
		)

		return
	}
	taskPatch := taskPatchFromRequest(request)

	taskDomain, err := h.tasksService.PatchTask(ctx, taskID, taskPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch task")
		return
	}
	response := PatchTaskReponse(TaskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)

}
func taskPatchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.TaskPatch{
		Title:       request.Title.ToDomain(),
		Description: request.Description.ToDomain(),
		Completed:   request.Comleted.ToDomain(),
	}
}

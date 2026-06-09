package users_transport_http

import (
	"net/http"

	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_request "github.com/Ravenmax/ToDo/internal/core/transport/http/request"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
)

type GetUserReponse UserDTOResponce

// GetUser       godoc
// @Summary      Получение пользователя
// @Description  Получения пользователя по ID
// @Tags         Users
// @Produce      json
// @Param        id path int true "ID пользователя"
// @Success      200  {object}  GetUserReponse "Найденный пользователь"
// @Failure      400  {object}  core_http_response.ErrorResponse "Bad request"
// @Failure      404  {object}  core_http_response.ErrorResponse "User not found"
// @Failure      500  {object}  core_http_response.ErrorResponse "Internal server error"
// @Router		 /users/{id} [get]
func (h *UsersHTTPHandlers) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	userID, err := core_http_request.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get UserID path Value")
		return
	}

	user, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get UserID")
		return
	}

	response := GetUserReponse(UserDTOFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}

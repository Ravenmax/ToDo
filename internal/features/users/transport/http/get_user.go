package users_transport_http

import (
	"net/http"

	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
	core_http_utils "github.com/Ravenmax/ToDo/internal/core/transport/http/utils"
)

type GetUserReponse UserDTOResponce

func (h *UsersHTTPHandlers) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	userID, err := core_http_utils.GetIntPathValues(r, "id")
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

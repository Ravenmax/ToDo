package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
	core_http_utils "github.com/Ravenmax/ToDo/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTOResponce

func (h *UsersHTTPHandlers) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offest, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get 'limit'/'offset' param")
		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offest)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to getUsers")

		return
	}
	response := GetUsersResponse(UsersDTOFromDomains(userDomains))
	responseHandler.JSONResponse(response, http.StatusOK)

}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}
	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}
	return limit, offset, nil
}

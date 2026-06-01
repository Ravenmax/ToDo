package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_request "github.com/Ravenmax/ToDo/internal/core/transport/http/request"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
)

type GetUsersResponse []UserDTOResponce

type UsersQueryParams struct {
	limit  *int
	offset *int
}

func (h *UsersHTTPHandlers) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	queryParams, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get 'limit'/'offset' param")
		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, queryParams.limit, queryParams.offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to getUsers")

		return
	}
	response := GetUsersResponse(UsersDTOFromDomains(userDomains))
	responseHandler.JSONResponse(response, http.StatusOK)

}

func getLimitOffsetQueryParams(r *http.Request) (UsersQueryParams, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)
	var (
		queryParams UsersQueryParams
		err         error
	)
	queryParams.limit, err = core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return UsersQueryParams{}, fmt.Errorf("get 'limit' query param: %w", err)
	}
	queryParams.offset, err = core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return UsersQueryParams{}, fmt.Errorf("get 'offset' query param: %w", err)
	}
	return queryParams, nil
}

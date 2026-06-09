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

// GetUsers  godoc
// @Summary      Список пользователей
// @Description  Просмотр списка пользователей с опциональными параметрами
// @Tags         Users
// @Produce      json
// @Param        limit query int false "Размер страницы с пользователями"
// @Param        offset query int false "Смещение страницы с пользователями"
// @Success      200  {object}  GetUsersResponse "Успешное получение списка пользователей"
// @Failure      400  {object}  core_http_response.ErrorResponse "Bad request"
// @Failure      404  {object}  core_http_response.ErrorResponse "User not found"
// @Failure      500  {object}  core_http_response.ErrorResponse "Internal server error"
// @Router		 /users [get]
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

package users_transport_http

import (
	"net/http"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_request "github.com/Ravenmax/ToDo/internal/core/transport/http/request"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"      validate:"required,min=3,max=100"                    example:"Ivan Ivanov"`
	PhoneNumber *string `json:"phone_number"   validate:"omitempty,min=10,max=100,startswith=+"     example:"+73336669999"` //можно использовать e164 в валидации номера телефона
}

type CreateUserResponse UserDTOResponce

// CreateUser  godoc
// @Summary      Создание пользователя
// @Description  Создание пользователя в системе из json body запроса
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body CreateUserRequest true "CreateUser тело запроса"
// @Success      201  {object}  CreateUserResponse "Успешно созданный пользователь"
// @Failure      400  {object}  core_http_response.ErrorResponse "Bad request"
// @Failure      500  {object}  core_http_response.ErrorResponse "Internal server error"
// @Router		 /users [post]
func (h *UsersHTTPHandlers) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}
	userDomain := domainFromDTO(request)
	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")

		return
	}

	respone := CreateUserResponse(UserDTOFromDomain(userDomain))

	responseHandler.JSONResponse(respone, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}

package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_request "github.com/Ravenmax/ToDo/internal/core/transport/http/request"
	core_http_response "github.com/Ravenmax/ToDo/internal/core/transport/http/response"
	core_http_types "github.com/Ravenmax/ToDo/internal/core/transport/http/types"
	"go.uber.org/zap"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"      swaggertype:"string" example:"Иван Иванович"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"   swaggertype:"string" example:"+73336669999"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("`FullName` cant be NULL")
		}
		fullNameLen := len([]rune(*r.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("`FullName` must be beetwen 3 and 100 symbols ")
		}
	}
	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("`PhoneNumber` must be beetwen 10 and 15 symbols ")
			}
			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("`PhoneNumber` must start from `+`")
			}
		}

	}
	return nil
}

type PatchUserResponse UserDTOResponce

// PatchUser  godoc
// @Summary      Обновление пользователя
// @Description  Обновления пользователя с заданным ID по JSON из тела запроса
// @Description  ### Логика обновления полей (Three-state logic):
// @Description  1.**Поле явно не передано**: `"phone_number"`, значение в БД не меняется.
// @Description  1.**Явно передано значение**: `"phone_number":"+73336669999"`, устанавливаем новый номер телефона.
// @Description  1.**Поле явно не передано**: `"phone_number":null`, очищаем поле в БД.
// @Description  Ограничения: `full_name` не может быть null.
// @Tags         Users
// @Accept		 json
// @Produce      json
// @Param        id path int true "ID пользователя"
// @Param        request body PatchUserRequest true "PatchUserRequest тело запроса"
// @Success      200  {object}  PatchUserResponse "Успешное обновление пользователя"
// @Failure      400  {object}  core_http_response.ErrorResponse "Bad request"
// @Failure      404  {object}  core_http_response.ErrorResponse "User not found"
// @Failure      409  {object}  core_http_response.ErrorResponse "Conflict"
// @Failure      500  {object}  core_http_response.ErrorResponse "Internal server error"
// @Router		 /users/{id} [patch]
func (h *UsersHTTPHandlers) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	userID, err := core_http_request.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user ID",
		)
		return
	}
	log.Debug("request", zap.String("URLPath", r.URL.Path))
	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode http request",
		)
		return
	}

	userPatch := UserPatchFromRequest(request)
	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user")
		return
	}

	response := PatchUserResponse(UserDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)

}

func UserPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.FullName.ToDomain(),
		request.PhoneNumber.ToDomain(),
	)

}

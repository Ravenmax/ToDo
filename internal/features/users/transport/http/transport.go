package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_http_server "github.com/Ravenmax/ToDo/internal/core/transport/http/server"
	"github.com/google/uuid"
)

type UsersHTTPHandlers struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		fullName string,
		phoneNumber *string,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		id uuid.UUID,
	) (domain.User, error)
	DeleteUser(
		ctx context.Context,
		id uuid.UUID,
	) error
	PatchUser(
		ctx context.Context,
		id uuid.UUID,
		patch domain.UserPatch,
	) (domain.User, error)
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandlers {
	return &UsersHTTPHandlers{
		usersService: usersService,
	}
}
func (h *UsersHTTPHandlers) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
			// Middleware: []core_http_middleware.Middleware{
			// 	core_http_middleware.Dummy("GetUsers middleware"),
			//
			// },
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
		},
	}
}

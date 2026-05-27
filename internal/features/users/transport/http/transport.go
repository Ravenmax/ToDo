package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_http_server "github.com/Ravenmax/ToDo/internal/core/transport/http/server"
)

type UsersHTTPHandlers struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)
	DeleteUser(
		ctx context.Context,
		id int,
	) error
	PatchUser(
		ctx context.Context,
		id int,
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
			Mehtod:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Mehtod:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
		{
			Mehtod:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Mehtod:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Mehtod:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
		},
	}
}

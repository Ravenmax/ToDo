package users_transport_http

import (
	"net/http"

	core_http_server "github.com/Ravenmax/ToDo/internal/core/transport/http/server"
)

type UsersHTTPHandlers struct {
	usersService UsersService
}

type UsersService interface {
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
	}
}

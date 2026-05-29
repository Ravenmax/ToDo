package users_transport_http

import "github.com/Ravenmax/ToDo/internal/core/domain"

type UserDTOResponce struct {
	ID          int     `json:"id"`
	Version     int64   `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func UserDTOFromDomain(user domain.User) UserDTOResponce {
	return UserDTOResponce{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}
func UsersDTOFromDomains(users []domain.User) []UserDTOResponce {
	usersDTO := make([]UserDTOResponce, len(users))
	for i, user := range users {
		usersDTO[i] = UserDTOFromDomain(user)
	}
	return usersDTO
}

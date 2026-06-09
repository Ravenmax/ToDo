package users_transport_http

import "github.com/Ravenmax/ToDo/internal/core/domain"

type UserDTOResponce struct {
	ID          int     `json:"id"              example:"10"`
	Version     int64   `json:"version"         example:"3"`
	FullName    string  `json:"full_name"       example:"Ivanov Ivan"`
	PhoneNumber *string `json:"phone_number"    example:"+73336669999"`
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

package users_postgres_repository

import (
	"github.com/Ravenmax/ToDo/internal/core/domain"
	"github.com/google/uuid"
)

type UserModel struct {
	ID          uuid.UUID `db:"id"`
	Version     int       `db:"version"`
	FullName    string    `db:"full_name"`
	PhoneNumber *string   `db:"phone_number"` // или string, но с NULL обработкой
}

func UsersDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))
	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.FullName,
			user.PhoneNumber,
		)
	}
	return userDomains
}

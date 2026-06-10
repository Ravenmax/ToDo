package users_service

import (
	"context"
	"fmt"

	"github.com/Ravenmax/ToDo/internal/core/domain"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	fullName string,
	phoneNumber *string,
) (domain.User, error) {
	user := domain.CreateUser(
		fullName,
		phoneNumber,
	)

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %w", err)
	}

	user, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}
	return user, nil
}

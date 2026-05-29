package users_service

import (
	"context"
	"fmt"

	"github.com/Ravenmax/ToDo/internal/core/domain"
)

func (s *UsersService) PatchUser(
	ctx context.Context,
	id int,
	patch domain.UserPatch,
) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}
	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("patch user: %w", err)
	}
	userPatched, err := s.usersRepository.PatchUser(ctx, user.ID, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("patch user: %w", err)
	}

	return userPatched, nil
}

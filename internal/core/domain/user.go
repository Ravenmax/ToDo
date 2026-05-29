package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/Ravenmax/ToDo/internal/core/errors"
)

type User struct {
	ID      int
	Version int64

	FullName    string
	PhoneNumber *string
}

func NewUser(
	id int,
	version int64,
	fullName string,
	phoneNumber *string,
) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func NewUserPatch(
	fullname Nullable[string],
	phoneNumber Nullable[string],
) UserPatch {
	return UserPatch{
		FullName:    fullname,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(fullName string, phoneNumber *string) User {
	return NewUser(UninitializedID, UninitializedVersion, fullName, phoneNumber)
}

func (u *User) Validate() error {
	fullnameLength := len([]rune(u.FullName))
	if fullnameLength < 3 || fullnameLength > 100 {
		return fmt.Errorf(
			"`FullName` must be beetwen 3 and 100 symbols: %d: %w",
			fullnameLength,
			core_errors.ErrInvalidArgument,
		)
	}
	if u.PhoneNumber != nil {
		phoneNumberLenght := len([]rune(*u.PhoneNumber))
		if phoneNumberLenght < 10 || phoneNumberLenght > 15 {
			return fmt.Errorf(
				"`PhoneNumber` must be beetwen 10 and 15 symbols: %d: %w",
				phoneNumberLenght,
				core_errors.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid PhoneNumber format: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}
	return nil
}

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func (p *UserPatch) Validate() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf(
			"FullName cant be patched to Null: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}
func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf(
			"validate user patch: %w",
			err,
		)
	}

	tmp := *u

	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}
	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}
	if err := tmp.Validate(); err != nil {
		return fmt.Errorf(
			"validate patched user: %w",
			err,
		)
	}

	*u = tmp

	return nil
}

package user

import (
	"CRUD/internal/domain/user"
)

type UpdateUserUseCase struct {
	repo user.Repository
}

func NewUpdateUserUseCase(repo user.Repository) *UpdateUserUseCase {
	return &UpdateUserUseCase{repo: repo}
}

func (uc *UpdateUserUseCase) Execute(id, name, email string) (*user.User, error) {
	existingUser, err := uc.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	existingUser.Name = name
	existingUser.Email = email

	if err := existingUser.Validate(); err != nil {
		return nil, err
	}

	if err := uc.repo.UpdateUser(existingUser); err != nil {
		return nil, err
	}

	return existingUser, nil
}

package user

import (
	"CRUD/internal/domain/user"
)

type CreateUserUseCase struct {
	repo user.Repository
}

func NewCreateUserUseCase(repo user.Repository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) Execute(id int, name, email string) (*user.User, error) {
	user, err := user.NewUser(id, name, email)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil

}

package user

import (
	"CRUD/internal/domain/user"
	"log"
)

type CreateUserUseCase struct {
	repo user.Repository
}

func NewCreateUserUseCase(repo user.Repository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) Execute(id, name, email string) (*user.User, error) {
	newUser, err := user.NewUser(id, name, email)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.CreateUser(newUser); err != nil {
		return nil, err
	}

	log.Printf("El usuario fue creado correctamente\n", newUser)
	return newUser, nil

}

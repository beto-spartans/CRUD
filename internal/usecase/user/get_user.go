package user

import (
	"CRUD/internal/domain/user"
	"log"
)

type GetUserUseCase struct {
	repo user.Repository
}

func NewGetUserUseCase(repo user.Repository) *GetUserUseCase {
	return &GetUserUseCase{repo: repo}
}

func (uc *GetUserUseCase) Execute(id string) (*user.User, error) {
	user, err := uc.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	log.Println("El usuario fue entregado correctamente")
	return user, nil
}

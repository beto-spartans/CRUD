package user

import (
	"CRUD/internal/domain/user"
	"log"
)

type DeleteUserUseCase struct {
	repo user.Repository
}

func NewDeleteUserUseCase(repo user.Repository) *DeleteUserUseCase {
	return &DeleteUserUseCase{repo: repo}
}

func (uc *DeleteUserUseCase) Delete(id string) error {
	if err := uc.repo.DeleteUser(id); err != nil {
		return err
	}
	log.Println("El usuario fue borrado correctamente")
	return nil

}

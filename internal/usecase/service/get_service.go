package service

import (
	"CRUD/internal/domain/service"
	//"database/sql"
	//"fmt"
	//"log"
	//"strconv"
)

type GetServicesUseCase struct {
	repo service.Repository
}

type Repository interface {
	GetServices(page, size int, queryFilter string) ([]service.Service, error) // Método para obtener todos los servicios
	// CreateUser(user *User) error
	// UpdateUser(user *User) error
	// Delete(id int) error
	// GetId(id int) error
	// GetAllUsers() ([]*User, error)
}

func NewCreateUserUseCase(repo service.Repository) *GetServicesUseCase {
	return &GetServicesUseCase{repo: repo}
}

func (uc *GetServicesUseCase) Execute(page, size int, queryFilter string) ([]service.Service, error) {
	return uc.repo.GetServices(page, size, queryFilter)
}

package user

import (
	"CRUD/internal/domain/user"
)

type UserServiceInterface interface {
	CreateUser(id, name, email string) (*user.User, error)
	GetUser(id string) (*user.User, error)
	UpdateUser(id, name, email string) (*user.User, error)
	DeleteUser(id string) error
}

type UserService struct {
	createUseCase *CreateUserUseCase
	getUseCase    *GetUserUseCase
	updateUseCase *UpdateUserUseCase
	deleteUseCase *DeleteUserUseCase
}

func NewUserService(
	create *CreateUserUseCase,
	get *GetUserUseCase,
	update *UpdateUserUseCase,
	delete *DeleteUserUseCase,
) *UserService {
	return &UserService{
		createUseCase: create,
		getUseCase:    get,
		updateUseCase: update,
		deleteUseCase: delete,
	}
}

func (s *UserService) CreateUser(id, name, email string) (*user.User, error) {
	return s.createUseCase.Execute(id, name, email)
}

func (s *UserService) GetUser(id string) (*user.User, error) {
	return s.getUseCase.Execute(id)
}

func (s *UserService) UpdateUser(id string, name, email string) (*user.User, error) {
	return s.updateUseCase.Execute(id, name, email)
}

func (s *UserService) DeleteUser(id string) error {
	return s.deleteUseCase.Delete(id)
}

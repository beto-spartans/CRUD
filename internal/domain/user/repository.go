package user

type Repository interface {
	CreateUser(user *User) error
	GetUserById(id string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

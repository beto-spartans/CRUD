package user

type Repository interface {
	CreateUser(user *User) error
	GetUserById(id int) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error
	GetAllUsers() ([]User, error)
}

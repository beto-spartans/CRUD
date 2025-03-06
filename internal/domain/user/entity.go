package user

import "errors"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(id, name, email string) (*User, error) {

	if email == "" {
		return nil, errors.New("email is required")
	}

	if name == "" {
		return nil, errors.New("name is required")
	}

	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

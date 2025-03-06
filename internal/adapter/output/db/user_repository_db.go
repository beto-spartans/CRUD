package db

import (
	"CRUD/internal/domain/user"
	"database/sql"
	"fmt"
	"log"
)

type UserRepositoryDB struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &UserRepositoryDB{db: db}
}

func (repo *UserRepositoryDB) CreateUser(u *user.User) error {
	query := `INSERT INTO users.users_table (id_user , name , email) VALUES ($1, $2, $3)`
	_, err := repo.db.Exec(query, u.ID, u.Name, u.Email)
	if err != nil {
		return fmt.Errorf("error al crear usuario: %w", err)
	}
	return nil
}

func (repo *UserRepositoryDB) UpdateUser(u *user.User) error {
	query := `UPDATE users SET name = $1, email = $2, age = $3 WHERE id = $4`
	result, err := repo.db.Exec(query, u.Name, u.Email, u.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar usuario: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error obteniendo filas afectadas: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró usuario con id %d", u.ID)
	}
	log.Printf("User %v actualizado", u.ID)
	return nil
}

func (repo *UserRepositoryDB) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error obteniendo filas afectadas: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró usuario con id %d", id)
	}
	log.Println("User eliminado")
	return nil
}

func (repo *UserRepositoryDB) GetUserById(id string) (*user.User, error) {
	query := `SELECT id_user, name, email FROM users.users_table WHERE id_user = $1`
	u := &user.User{}
	err := repo.db.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuario: %w", err)
	}
	return u, nil
}

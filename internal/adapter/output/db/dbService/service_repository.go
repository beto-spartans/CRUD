package dbservice

import (
	"CRUD/internal/domain/service"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type ServiceRepository struct {
	db *sql.DB
}
type Repository interface {
	GetServices(page, size int, queryFilter string) ([]service.Service, error) // Método para obtener todos los servicios
	// cambio
}

func NewServiceRepository(db *sql.DB) service.Repository {
	return &ServiceRepository{db: db}
}

func (repo *ServiceRepository) GetServices(page, size int, queryFilter string) ([]service.Service, error) {
	start := (page - 1) * size
	baseQuery := `SELECT id, name, version, added_by, release_date, assigned, status FROM services`
	whereClause := ""
	var params []interface{}
	paramIndex := 1

	if queryFilter != "" {
		whereClause = `WHERE name ILIKE $` + strconv.Itoa(paramIndex) +
			`OR version ILIKE $` + strconv.Itoa(paramIndex) +
			`OR added_by ILIKE $` + strconv.Itoa(paramIndex)
		params = append(params, "%"+queryFilter+"%")
		paramIndex++
	}
	// Agregar limitación y offset
	limitOffsetClause := ` LIMIT $` + strconv.Itoa(paramIndex) + ` OFFSET $` + strconv.Itoa(paramIndex+1)
	params = append(params, size, start)

	query := baseQuery + whereClause + limitOffsetClause

	// Logging para depuración
	log.Println("Executing query:", query)
	log.Println("With parameters:", params)

	// Ejecutar consulta
	rows, err := repo.db.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando consulta: %w", err)
	}
	defer rows.Close()

	var services []service.Service
	for rows.Next() {
		var s service.Service
		if err := rows.Scan(&s.ID, &s.Name, &s.Version, &s.AddedBy, &s.ReleaseDate, &s.Assigned, &s.Status); err != nil {
			return nil, fmt.Errorf("error escaneando fila: %w", err)
		}
		services = append(services, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando filas: %w", err)
	}
	return services, nil

}

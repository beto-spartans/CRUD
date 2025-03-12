package service

type Repository interface {
	GetServices(page, size int, queryFilter string) ([]Service, error)
}

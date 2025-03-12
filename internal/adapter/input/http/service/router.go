package service

import (
	"CRUD/internal/usecase/service"
	"net/http"
)

func NewRouter(getServicesUC *service.GetServicesUseCase) http.Handler {
	controller := NewServiceController(getServicesUC)
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/cartridges/service", controller.GetServicesHandler)
	return mux
}

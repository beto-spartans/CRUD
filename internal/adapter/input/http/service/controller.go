package service

import (
	"CRUD/internal/usecase/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type ServiceControler struct {
	getService *service.GetServicesUseCase
}

func NewServiceController(getServicesUC *service.GetServicesUseCase) *ServiceControler {
	return &ServiceControler{getService: getServicesUC}
}

func (c *ServiceControler) GetServicesHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	queryFilter := r.URL.Query().Get("queryFilter")

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	services, err := c.getService.Execute(page, size, queryFilter)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := map[string]any{"services": services}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

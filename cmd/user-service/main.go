package main

import (
	"fmt"
	"log"
	"net/http"

	httpAdapter "CRUD/internal/adapter/input/http"
	dbAdapter "CRUD/internal/adapter/output/db"
	"CRUD/internal/config"
	"CRUD/internal/usecase/user"
)

func main() {
	cfg := config.NewConfig()
	dataSourceName := cfg.DataSourceName()

	db, err := dbAdapter.InitDB(dataSourceName)
	if err != nil {
		log.Fatalf("error inicializando la base de datos: %v", err)
	}

	userRepo := dbAdapter.NewUserRepository(db)

	createUC := user.NewCreateUserUseCase(userRepo)
	getUC := user.NewGetUserUseCase(userRepo)
	updateUC := user.NewUpdateUserUseCase(userRepo)
	deleteUC := user.NewDeleteUserUseCase(userRepo)

	userService := user.NewUserService(createUC, getUC, updateUC, deleteUC)
	userController := httpAdapter.NewUserController(userService)
	router := httpAdapter.NewRouter(userController)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

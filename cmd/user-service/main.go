package main

import (
	"fmt"
	"log"
	"net/http"

	dbAdapter "CRUD/internal/adapter/output/db"
	dbAdapterServ "CRUD/internal/adapter/output/db/dbService"
	"CRUD/internal/config"
	"CRUD/internal/usecase/service"

	//"CRUD/internal/adapter/http"
	v1 "CRUD/internal/adapter/input/http/service"
)

func main() {
	cfg := config.NewConfig()
	dataSourceName := cfg.DataSourceName()

	db, err := dbAdapter.InitDB(dataSourceName)
	if err != nil {
		log.Fatalf("error inicializando la base de datos: %v", err)
	}

	//userRepo := dbAdapter.NewUserRepository(db)
	serviceRepo := dbAdapterServ.NewServiceRepository(db)
	getService := service.NewCreateUserUseCase(serviceRepo)
	// createUC := user.NewCreateUserUseCase(userRepo)
	// getUC := user.NewGetUserUseCase(userRepo)
	// updateUC := user.NewUpdateUserUseCase(userRepo)
	// deleteUC := user.NewDeleteUserUseCase(userRepo)

	//userService := user.NewUserService(createUC, getUC, updateUC, deleteUC)
	//userController := httpAdapter.NewUserController(userService)
	//router := httpAdapter.NewRouter(userController)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	router := v1.NewRouter(getService)

	log.Fatal(http.ListenAndServe(":8080", router))

}

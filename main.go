package main

import (
	"go-mysql-api/pkg/infrastructure/config"
	"go-mysql-api/pkg/infrastructure/database"
	"go-mysql-api/pkg/infrastructure/server"
	"go-mysql-api/pkg/user/handler"
	"go-mysql-api/pkg/user/repository"
	"go-mysql-api/pkg/user/usecase"
	"go-mysql-api/pkg/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	utils.LoggingSetting()
	cfg := config.NewConfig()
	server := server.NewServer(cfg)
	db := database.NewDB(cfg).Connect()
	router := server.Router

	// User
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	handler.NewUserHandler(router, userUsecase)

	server.Run()
}

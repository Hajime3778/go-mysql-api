package main

import (
	"go-mysql-api/pkg/infrastructure"
	"go-mysql-api/pkg/infrastructure/database"
	"go-mysql-api/pkg/user/handler"
	"go-mysql-api/pkg/user/repository"
	"go-mysql-api/pkg/user/usecase"
	"go-mysql-api/pkg/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	utils.LoggingSetting()
	router := infrastructure.NewRouting()
	db := database.NewDB()

	userRepo := repository.NewUserRepository(db.Connect())
	userUsecase := usecase.NewUserUsecase(userRepo)
	handler.NewUserHandler(router, userUsecase)

	router.Run(":3000")
}

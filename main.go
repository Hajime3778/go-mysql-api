package main

import (
	"go-mysql-api/pkg/infrastructure/database"
	"go-mysql-api/pkg/user/handler"
	"go-mysql-api/pkg/user/repository"
	"go-mysql-api/pkg/user/usecase"
	"go-mysql-api/pkg/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	utils.LoggingSettings()

	router := gin.Default()

	// すべてのアクセス許可
	config := cors.Config{AllowOrigins: []string{"*"}}
	router.Use(cors.New(config))

	router.StaticFile("/", "./index.html")

	db := database.NewDB()
	userRepo := repository.NewUserRepository(db.Connect())
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// User routing set up
	router.GET("api/user", userHandler.GetAll)
	router.GET("api/user/:id", userHandler.Get)
	router.POST("api/user", userHandler.Create)
	router.PUT("api/user", userHandler.Update)
	router.DELETE("api/user/:id", userHandler.Delete)

	// POSTで更新したい場合↓のように書ける
	// router.POST("/user/*action", user.UpdateUser)

	router.Run(":3000")
}

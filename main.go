package main

import (
	"go-mysql-api/pkg/controllers"
	"go-mysql-api/pkg/infrastructure/database"
	"go-mysql-api/pkg/repositories"
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
	userRepo := repositories.NewUserRepository(db.Connect())
	userController := controllers.NewUserController(userRepo)

	// User routing set up
	router.GET("api/user", userController.GetUsers)
	router.GET("api/user/:id", userController.GetUser)
	router.POST("api/user", userController.CreateUser)
	router.PUT("api/user", userController.UpdateUser)
	router.DELETE("api/user/:id", userController.DeleteUser)

	// POSTで更新したい場合↓のように書ける
	// router.POST("/user/*action", user.UpdateUser)

	router.Run(":3000")
}

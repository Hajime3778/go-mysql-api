package server

import (
	"go-mysql-api/pkg/user/handler"
	"go-mysql-api/pkg/user/repository"
	"go-mysql-api/pkg/user/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	router := gin.Default()

	// すべてのアクセス許可
	config := cors.Config{AllowOrigins: []string{"*"}}
	router.Use(cors.New(config))
	router.StaticFile("/", "./index.html")

	return router
}

// SetUpRouter Setup all api routing
func (s *Server) SetUpRouter() {
	// Group v1
	apiV1 := s.router.Group("api/v1")
	s.userRoutes(apiV1)
}

func (s *Server) userRoutes(api *gin.RouterGroup) {

	repository := repository.NewUserRepository(s.db)
	usecase := usecase.NewUserUsecase(repository)
	handler := handler.NewUserHandler(s.router, usecase)

	userRoutes := api.Group("/users")
	{
		userRoutes.GET("", handler.GetAll)
		userRoutes.GET("/:id", handler.FindByID)
		userRoutes.POST("", handler.Create)
		userRoutes.PUT("", handler.Update)
		userRoutes.DELETE("/:id", handler.Delete)
	}
}

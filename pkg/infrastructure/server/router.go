package server

import (
	"go-mysql-api/pkg/user/handler"
	"go-mysql-api/pkg/user/repository"
	"go-mysql-api/pkg/user/usecase"

	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	router := gin.Default()

	// すべてのアクセス許可
	// config := cors.Config{AllowOrigins: []string{"*"}}
	// router.Use(cors.New(config))
	// router.StaticFile("/", "./index.html")

	return router
}

// SetUpRouter Setup all api routing
func (s *Server) SetUpRouter() {

	// Group : v1
	//apiV1 := s.router.Group("api/v1")
	//s.userRoutes(apiV1)

	// apiV1s := "api/v1"
	// s.userRoutes2(apiV1s)

	repository := repository.NewUserRepository(s.db)
	usecase := usecase.NewUserUsecase(repository)
	handler := handler.NewUserHandler(s.router, usecase)
	api := s.router.Group("api/v1/users")
	{
		api.GET("", handler.GetAll)
		api.GET("/:id", handler.FindByID)
		api.POST("", handler.Create)
		api.PUT("", handler.Update)
		api.DELETE("/:id", handler.Delete)
	}

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

func (s *Server) userRoutes2(path string) {
	repository := repository.NewUserRepository(s.db)
	usecase := usecase.NewUserUsecase(repository)
	handler := handler.NewUserHandler(s.router, usecase)

	s.router.GET(path+"user", handler.GetAll)
	s.router.GET(path+"user/:id", handler.FindByID)
	s.router.POST(path+"user", handler.Create)
	s.router.PUT(path+"user", handler.Update)
	s.router.DELETE(path+"user/:id", handler.Delete)
}

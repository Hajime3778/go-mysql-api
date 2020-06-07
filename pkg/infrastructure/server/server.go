package server

import (
	"go-mysql-api/pkg/infrastructure/config"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server server
type Server struct {
	Router *gin.Engine
	server *http.Server
}

// NewServer Server create
func NewServer(c *config.Config) *Server {
	r := newRouter()
	s := newServer(c, r)
	return &Server{
		Router: r,
		server: s,
	}
}

func newRouter() *gin.Engine {
	router := gin.Default()

	// すべてのアクセス許可
	config := cors.Config{AllowOrigins: []string{"*"}}
	router.Use(cors.New(config))
	router.StaticFile("/", "./index.html")

	return router
}

func newServer(c *config.Config, router *gin.Engine) *http.Server {
	s := &http.Server{
		Addr:         c.Server.Port,
		Handler:      router,
		ReadTimeout:  time.Duration(c.Server.Timeout) * time.Second,
		WriteTimeout: time.Duration(c.Server.Timeout) * time.Second,
	}
	return s
}

// Run Run server
func (s *Server) Run() {
	s.server.ListenAndServe()
}

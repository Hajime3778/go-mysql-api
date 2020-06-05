package infrastructure

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouting foo
func NewRouting() *gin.Engine {
	router := gin.Default()

	// すべてのアクセス許可
	config := cors.Config{AllowOrigins: []string{"*"}}
	router.Use(cors.New(config))
	router.StaticFile("/", "./index.html")

	return router
}

package main_test

import (
	"go-mysql-api/pkg/controllers"
	"go-mysql-api/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestMainRouter(t *testing.T) {

	utils.LoggingSettings()

	router := gin.Default()

	// すべてのアクセス許可
	config := cors.Config{AllowOrigins: []string{"*"}}
	router.Use(cors.New(config))

	user := controllers.NewUserController()
	router.GET("api/user", user.GetUsers)
	//router.Run(":3000")

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

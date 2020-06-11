package main_test

import (
	"go-mysql-api/pkg/infrastructure/config"
	"go-mysql-api/pkg/infrastructure/database"
	"go-mysql-api/pkg/infrastructure/server"
	"go-mysql-api/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	utils.LoggingSetting()
	cfg := config.NewConfig()
	db := database.NewDB(cfg)

	server := server.NewServer(cfg, db)
	router := server.SetUpRouter()

	testUsers(t, router)
}

func testUsers(t *testing.T, router *gin.Engine) {
	getAllRes := httptest.NewRecorder()
	getAllReq, _ := http.NewRequest("GET", "/api/v1/users", nil)
	router.ServeHTTP(getAllRes, getAllReq)

	assert.Equal(t, 200, getAllRes.Code)

	findByIdRes := httptest.NewRecorder()
	findByIdReq, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
	router.ServeHTTP(findByIdRes, findByIdReq)

	assert.Equal(t, 200, findByIdRes.Code)
}

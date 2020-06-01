package controllers_test

import (
	"go-mysql-api/pkg/controllers"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestUserController_GetUsers(t *testing.T) {
	// リクエスト生成
	req, _ : = http.NewRequest("GET", "/api/user", nil)

	// Contextセット
	var context *gin.Context
	context = &gin.Context{Request: req}

	user := controllers.NewUserController()
	user.GetUsers(context)
}

package controllers

import (
	"log"
	"net/http"
	"strconv"

	"go-mysql-api/pkg/models"
	"go-mysql-api/pkg/repositories"

	"github.com/gin-gonic/gin"
)

// UserController controller for user request
type UserController struct{}

// NewUserController is init for UserController
func NewUserController() *UserController {
	return new(UserController)
}

// GetUsers 複数のUserを取得します
func (u *UserController) GetUsers(c *gin.Context) {
	userRepository := repositories.NewUserRepository()

	result, err := userRepository.GetAll()
	if err != nil {
		log.Panicln(err)
	}

	c.JSON(http.StatusOK, result)
}

// GetUser 1件のUserを取得します
func (u *UserController) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userRepository := repositories.NewUserRepository()

	result, err := userRepository.FindByID(id)
	if err != nil {
		log.Panicln(err)
	}

	c.JSON(http.StatusOK, result)
}

// CreateUser Userを作成します
func (u *UserController) CreateUser(c *gin.Context) {
	userRepository := repositories.NewUserRepository()
	var user models.User
	c.BindJSON(&user)

	err := userRepository.Regist(user)
	if err != nil {
		log.Panicln(err)
	}
	c.JSON(http.StatusCreated, nil)
}

// UpdateUser Userを更新します。
func (u *UserController) UpdateUser(c *gin.Context) {
	userRepository := repositories.NewUserRepository()
	var user models.User
	c.BindJSON(&user)

	err := userRepository.Update(user)
	if err != nil {
		log.Panicln(err)
	}
	c.JSON(http.StatusOK, nil)
}

// DeleteUser Userを削除します
func (u *UserController) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userRepository := repositories.NewUserRepository()

	err := userRepository.Delete(id)
	if err != nil {
		log.Panicln(err)
	}
	c.JSON(http.StatusNoContent, nil)
}

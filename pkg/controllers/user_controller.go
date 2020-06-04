package controllers

import (
	"log"
	"net/http"
	"strconv"

	"go-mysql-api/pkg/domain"
	"go-mysql-api/pkg/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserController controller for user request
type UserController interface {
	GetUsers(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

// userController controller for user request
type userController struct {
	repo repositories.UserRepository
}

// NewUserController is init for UserController
func NewUserController(repo repositories.UserRepository) UserController {
	return &userController{
		repo: repo,
	}
}

// GetUsers 複数のUserを取得します
func (u *userController) GetUsers(c *gin.Context) {
	result, err := u.repo.GetAll()

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetUser 1件のUserを取得します
func (u *userController) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := u.repo.FindByID(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// CreateUser Userを作成します
func (u *userController) CreateUser(c *gin.Context) {
	var user domain.User
	c.BindJSON(&user)

	err := u.repo.Regist(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

// UpdateUser Userを更新します。
func (u *userController) UpdateUser(c *gin.Context) {
	var user domain.User
	c.BindJSON(&user)

	err := u.repo.Update(user)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

// DeleteUser Userを削除します
func (u *userController) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := u.repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

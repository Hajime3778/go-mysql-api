package handler

import (
	"log"
	"net/http"
	"strconv"

	"go-mysql-api/pkg/domain"
	"go-mysql-api/pkg/user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserHandler controller for user request
type UserHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// userHandler controller for user request
type userHandler struct {
	usecase usecase.UserUsecase
}

// NewUserHandler is init for UserHandler
func NewUserHandler(router *gin.Engine, u usecase.UserUsecase) {
	handler := &userHandler{
		usecase: u,
	}
	router.GET("api/user", handler.GetAll)
	router.GET("api/user/:id", handler.Get)
	router.POST("api/user", handler.Create)
	router.PUT("api/user", handler.Update)
	router.DELETE("api/user/:id", handler.Delete)
	// POSTで更新したい場合↓のように書ける
	// router.POST("/user/*action", handler.Update)
}

// GetAll 複数のUserを取得します
func (h *userHandler) GetAll(c *gin.Context) {
	result, err := h.usecase.GetUsers()

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

// Get 1件のUserを取得します
func (h *userHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := h.usecase.GetUser(id)

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

// Create Userを作成します
func (h *userHandler) Create(c *gin.Context) {
	var user domain.User
	c.BindJSON(&user)

	err := h.usecase.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

// Update Userを更新します。
func (h *userHandler) Update(c *gin.Context) {
	var user domain.User
	c.BindJSON(&user)

	err := h.usecase.UpdateUser(user)

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

// Delete Userを削除します
func (h *userHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.usecase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

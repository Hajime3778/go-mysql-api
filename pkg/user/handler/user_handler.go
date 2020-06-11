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
	FindByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// UserHandler controller for user request
type userHandler struct {
	router  *gin.Engine
	usecase usecase.UserUsecase
}

// NewUserHandler is init for UserHandler
func NewUserHandler(r *gin.Engine, u usecase.UserUsecase) UserHandler {
	return &userHandler{
		router:  r,
		usecase: u,
	}
}

// GetAll 複数のUserを取得します
func (h *userHandler) GetAll(c *gin.Context) {
	result, err := h.usecase.GetAll()

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

// FindByID 1件のUserを取得します
func (h *userHandler) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := h.usecase.FindByID(id)

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

	err := h.usecase.Create(user)
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

	err := h.usecase.Update(user)

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

	err := h.usecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

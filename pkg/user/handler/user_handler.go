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
// type UserHandler interface {
// 	Handle()
// 	// GetAll(c *gin.Context)
// 	// Get(c *gin.Context)
// 	// Create(c *gin.Context)
// 	// Update(c *gin.Context)
// 	// Delete(c *gin.Context)
// }

// UserHandler controller for user request
type UserHandler struct {
	router  *gin.Engine
	usecase usecase.UserUsecase
}

// NewUserHandler is init for UserHandler
func NewUserHandler(r *gin.Engine, u usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		router:  r,
		usecase: u,
	}
}

// Handle Set api handling
func (h *UserHandler) Handle() {
	h.router.GET("api/user", h.getAll)
	h.router.GET("api/user/:id", h.get)
	h.router.POST("api/user", h.create)
	h.router.PUT("api/user", h.update)
	h.router.DELETE("api/user/:id", h.delete)
	// POSTで更新したい場合↓のように書ける
	// router.POST("/user/*action", handler.Update)
}

// GetAll 複数のUserを取得します
func (h *UserHandler) getAll(c *gin.Context) {
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
func (h *UserHandler) get(c *gin.Context) {
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
func (h *UserHandler) create(c *gin.Context) {
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
func (h *UserHandler) update(c *gin.Context) {
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
func (h *UserHandler) delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.usecase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

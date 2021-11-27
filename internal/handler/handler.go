package handler

import (
	"golang-shop/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService *service.UserService
}

func NewHandler() *Handler {
	return &Handler{
		UserService: service.NewUserService(),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.GET("/", h.getAll)
		user.POST("/", h.create)
		user.GET("/:id", h.getOne)
		user.PUT("/:id", h.update)
		user.DELETE("/:id", h.delete)
	}

	return router
}

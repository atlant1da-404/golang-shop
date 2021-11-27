package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) create(c *gin.Context) {
	id := h.UserService.Create(1)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": id,
		"context": "Test API",
	})
}

func (h *Handler) getAll(c *gin.Context) {
	id := h.UserService.GetAll()
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": id,
		"context": "Test API",
	})
}

func (h *Handler) getOne(c *gin.Context) {
	id := h.UserService.GetOne(1)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": id,
		"context": "Test API",
	})
}

func (h *Handler) delete(c *gin.Context) {
	id := h.UserService.Delete(1)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": id,
		"context": "Test API",
	})
}

func (h *Handler) update(c *gin.Context) {
	id := h.UserService.Update(1)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": id,
		"context": "Test API",
	})
}

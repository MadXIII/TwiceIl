package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Products(c *gin.Context) {
	products, status, err := h.service.ToGet()
	if err != nil {
		c.JSON(status, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": products,
	})
}

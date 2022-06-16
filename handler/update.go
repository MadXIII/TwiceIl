package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madxiii/twiceil/model"
)

func (h *Handler) Update(c *gin.Context) {
	var body model.Product
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error":  "invalid body",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

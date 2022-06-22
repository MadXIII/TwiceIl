package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madxiii/twiceil/model"
)

func (h *Handler) Find(c *gin.Context) {
	var body model.Search
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error":  "invalid body",
		})
		return
	}
	product, err := h.service.ToFind(body.SearchName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

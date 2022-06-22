package handler

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/madxiii/twiceil/model"
)

func (h *Handler) ParseCreate(c *gin.Context) {
	temp, err := template.ParseFiles("client/create.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}
	if err := temp.Execute(c.Writer, nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}
}

func (h *Handler) Create(c *gin.Context) {
	var body model.Product
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error":  "invalid body",
		})
		return
	}

	id, status, err := h.service.ToCreate(&body)
	if err != nil {
		c.JSON(status, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"productId": id,
	})
}

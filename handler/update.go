package handler

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madxiii/twiceil/model"
)

func (h *Handler) ParseUpdate(c *gin.Context) {
	temp, err := template.ParseFiles("client/edit.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	param := c.Param("id")

	product, status, err := h.service.ToProduct(param)
	if err != nil {
		c.JSON(status, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	if err = temp.Execute(c.Writer, product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}
}

func (h *Handler) Update(c *gin.Context) {
	var body model.Product
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error":  "invalid body",
		})
		return
	}
	status, err := h.service.ToUpdate(&body)
	if err != nil {
		c.JSON(status, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

package handler

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Products(c *gin.Context) {
	temp, err := template.ParseFiles("client/main.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})

		return
	}

	products, status, err := h.service.ToProducts()
	if err != nil {
		c.JSON(status, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	if err := temp.Execute(c.Writer, products); err != nil {
		c.JSON(status, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}
}
